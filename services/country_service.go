package services

import (
	"TravelSphere/models"
	"TravelSphere/utils"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

var (
	countriesCache     []models.Country
	countriesCacheTime time.Time
	cacheMutex         sync.RWMutex
	cacheDuration      = 10 * time.Minute
)

type CountryService struct {
	client           *utils.APIClient
	countriesBaseURL string
	tripBaseURL      string
	apiKey           string
}

// APIResponseWrapper matches the exact JSON envelope hierarchy: {"data": {"objects": [...]}}
type APIResponseWrapper struct {
	Data Container `json:"data"`
}

type Container struct {
	Objects []v5RawCountry `json:"objects"`
}

type v5RawCountry struct {
	Names struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	} `json:"names"`
	Region     string `json:"region"`
	Population int64  `json:"population"`
	Capitals   []struct {
		Name string `json:"name"`
	} `json:"capitals"`
	Flag struct {
		UrlPng string `json:"url_png"`
		UrlSvg string `json:"url_svg"`
	} `json:"flag"`
	Currencies []struct {
		Code string `json:"code"`
	} `json:"currencies"`
	Languages []struct {
		Name string `json:"name"`
	} `json:"languages"`
}

func NewCountryService(client *utils.APIClient, apiKey string) *CountryService {
	countriesURL := beego.AppConfig.DefaultString("countries_base_url", "")
	tripURL := beego.AppConfig.DefaultString("trip_base_url", "")

	return &CountryService{
		client:           client,
		countriesBaseURL: countriesURL,
		tripBaseURL:      tripURL,
		apiKey:           apiKey,
	}
}

func (s *CountryService) FetchAllCountries() ([]models.Country, error) {
	cacheMutex.RLock()
	if countriesCache != nil && time.Since(countriesCacheTime) < cacheDuration {
		defer cacheMutex.RUnlock()
		return countriesCache, nil
	}
	cacheMutex.RUnlock()

	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	if countriesCache != nil && time.Since(countriesCacheTime) < cacheDuration {
		return countriesCache, nil
	}

	url := strings.TrimSuffix(s.countriesBaseURL, "/") + "?pretty=1"
	logs.Info("Fetching countries from external API: %s", url)

	headers := map[string]string{
		"Authorization": "Bearer " + s.apiKey,
	}

	body, code, err := s.client.Get(url, headers)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch countries: %w", err)
	}
	if code != 200 {
		return nil, fmt.Errorf("REST Countries returned status code %d", code)
	}

	// Unmarshal using the nested wrapper structure matching the layout payload
	var wrapper APIResponseWrapper
	if err := json.Unmarshal(body, &wrapper); err != nil {
		return nil, fmt.Errorf("failed to decode nested layout payload: %w", err)
	}

	rawList := wrapper.Data.Objects

	if len(rawList) == 0 {
		logs.Error("CRITICAL: Mismatch structure. Objects array is empty.")
		return nil, fmt.Errorf("no country data found in objects list")
	}

	parsedList := make([]models.Country, 0)

	for _, raw := range rawList {
		commonName := raw.Names.Common
		if commonName == "" {
			continue
		}

		flag := raw.Flag.UrlPng
		if flag == "" {
			flag = raw.Flag.UrlSvg
		}

		capital := "N/A"
		if len(raw.Capitals) > 0 {
			capital = raw.Capitals[0].Name
		}

		var dynamicLanguages []string
		for _, lang := range raw.Languages {
			if lang.Name != "" {
				dynamicLanguages = append(dynamicLanguages, lang.Name)
			}
		}

		var dynamicCurrencies []string
		for _, curr := range raw.Currencies {
			if curr.Code != "" {
				dynamicCurrencies = append(dynamicCurrencies, curr.Code)
			}
		}
		formattedCurrencies := strings.Join(dynamicCurrencies, ", ")
		if formattedCurrencies == "" {
			formattedCurrencies = "N/A"
		}

		slug := strings.ReplaceAll(strings.ToLower(commonName), " ", "-")

		c := models.Country{
			Name:                raw.Names.Official,
			CommonName:          commonName,
			Slug:                slug,
			Flag:                flag,
			Capital:             capital,
			Region:              raw.Region,
			Population:          raw.Population,
			FormattedPopulation: utils.FormatPopulation(raw.Population),
			Currencies:          formattedCurrencies,
			Languages:           strings.Join(dynamicLanguages, ", "),
		}
		parsedList = append(parsedList, c)
	}

	countriesCache = parsedList
	countriesCacheTime = time.Now()
	logs.Info("Successfully cached %d countries from objects layout structure", len(parsedList))

	return parsedList, nil
}

func (s *CountryService) List(search, region string) ([]models.Country, error) {
	list, err := s.FetchAllCountries()
	if err != nil {
		return nil, err
	}

	var filtered []models.Country
	searchLower := strings.ToLower(strings.TrimSpace(search))
	regionLower := strings.ToLower(strings.TrimSpace(region))

	for _, c := range list {
		matchSearch := true
		matchRegion := true

		if searchLower != "" {
			matchSearch = strings.Contains(strings.ToLower(c.CommonName), searchLower) ||
				strings.Contains(strings.ToLower(c.Name), searchLower) ||
				strings.Contains(strings.ToLower(c.Slug), searchLower)
		}

		if regionLower != "" {
			matchRegion = strings.ToLower(c.Region) == regionLower
		}

		if matchSearch && matchRegion {
			filtered = append(filtered, c)
		}
	}

	return filtered, nil
}

func (s *CountryService) GetBySlug(slug string) (models.Country, error) {
	slugLower := strings.ToLower(strings.TrimSpace(slug))
	if slugLower == "" {
		return models.Country{}, fmt.Errorf("slug cannot be empty")
	}

	list, err := s.FetchAllCountries()
	if err != nil {
		return models.Country{}, err
	}

	for _, c := range list {
		if c.Slug == slugLower {
			return c, nil
		}
	}

	return models.Country{}, fmt.Errorf("country not found for slug: %s", slug)
}

// GetAttractions retrieves tourist attractions near specific geographical coordinates.
func (s *CountryService) GetAttractions(lat, lon float64) ([]models.Attraction, error) {
	apiKey := os.Getenv("OPENTRIPMAP_KEY")
	if apiKey == "" {
		// Graceful fallback with mock items if API key is not set
		return []models.Attraction{
			{Name: "Historic Center & Landmarks", Kinds: "monuments,historic", Dist: 120.5, Wikidata: "Q12345"},
			{Name: "National Museum of Art", Kinds: "museums,arts", Dist: 450.2, Wikidata: "Q67890"},
			{Name: "Central Botanical Park", Kinds: "nature,parks", Dist: 1100.8, Wikidata: "Q54321"},
		}, nil
	}

	// Build URL for places within 50km (50000 meters)
	queryURL := fmt.Sprintf("%s/en/places/radius?radius=50000&lon=%f&lat=%f&apikey=%s&limit=10&format=json",
		s.tripBaseURL, lon, lat, url.QueryEscape(apiKey))

	body, code, err := s.client.Get(queryURL, nil)
	if err != nil {
		// Graceful fallback on API timeout or network failure
		return []models.Attraction{
			{Name: "Local Attractions (Offline/Fallback)", Kinds: "tourist_attractions", Dist: 0.0},
		}, nil
	}

	if code != 200 {
		return []models.Attraction{
			{Name: "Local Attractions (Unavailable)", Kinds: "tourist_attractions", Dist: 0.0},
		}, nil
	}

	var attractions []models.Attraction
	if err := json.Unmarshal(body, &attractions); err != nil {
		return nil, fmt.Errorf("failed to parse attractions response: %w", err)
	}

	// Filter out attractions without a valid name
	var filtered []models.Attraction
	for _, attr := range attractions {
		if strings.TrimSpace(attr.Name) != "" {
			filtered = append(filtered, attr)
		}
	}

	if len(filtered) == 0 {
		return []models.Attraction{
			{Name: "Scenic Capital View", Kinds: "viewpoints", Dist: 100.0},
		}, nil
	}

	return filtered, nil
}
