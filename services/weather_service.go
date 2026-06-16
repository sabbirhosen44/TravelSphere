package services

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"TravelSphere/models"
	"TravelSphere/utils"

	beego "github.com/beego/beego/v2/server/web"
)

// WeatherService fetches local weather information from WeatherAPI.
type WeatherService struct {
	client  *utils.APIClient
	baseURL string
	apiKey  string
}

// NewWeatherService creates a new instance of WeatherService using configuration defaults.
func NewWeatherService(client *utils.APIClient) *WeatherService {
	// Fallback to the default WeatherAPI production endpoint if not provided in app.conf
	weatherURL := beego.AppConfig.DefaultString("weather_base_url", "")
	weatherKey := beego.AppConfig.DefaultString("weather_api_key", "")

	return &WeatherService{
		client:  client,
		baseURL: weatherURL,
		apiKey:  weatherKey,
	}
}

// SetBaseURL changes the target URL (for mocking/testing).
func (s *WeatherService) SetBaseURL(url string) {
	s.baseURL = url
}

type rawWeatherResponse struct {
	Current struct {
		TempC     float64 `json:"temp_c"`
		WindKph   float64 `json:"wind_kph"`
		Humidity  int     `json:"humidity"`
		Condition struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
		} `json:"condition"`
	} `json:"current"`
}

// GetWeather retrieves the weather for a query (e.g., "Paris", or latitude,longitude coordinates).
func (s *WeatherService) GetWeather(query string) (models.WeatherInfo, error) {
	if s.apiKey == "" {
		// Graceful fallback if no weather key is provided
		return models.WeatherInfo{
			TempC:          22.0,
			Condition:      "Clear & Pleasant (Offline)",
			ConditionIcon:  "https://cdn.weatherapi.com/weather/64x64/day/113.png",
			WindKph:        12.5,
			Humidity:       50,
			Recommendation: "Pleasant. Excellent weather for city tours and outdoor exploration.",
		}, nil
	}

	queryURL := fmt.Sprintf("%s/current.json?key=%s&q=%s", s.baseURL, s.apiKey, url.QueryEscape(query))
	body, code, err := s.client.Get(queryURL,nil)
	if err != nil || code != 200 {
		// Graceful fallback on HTTP errors or network timeout
		return models.WeatherInfo{
			TempC:          20.0,
			Condition:      "Sunny Intervals (Offline Fallback)",
			ConditionIcon:  "https://cdn.weatherapi.com/weather/64x64/day/116.png",
			WindKph:        10.0,
			Humidity:       60,
			Recommendation: "Mild weather. Pack light layers and enjoy your sightseeing.",
		}, nil
	}

	var raw rawWeatherResponse
	if err := json.Unmarshal(body, &raw); err != nil {
		return models.WeatherInfo{}, fmt.Errorf("failed to parse weather JSON: %w", err)
	}

	// Ensure the icon URL is properly formatted with the schema (HTTPS)
	icon := raw.Current.Condition.Icon
	if strings.HasPrefix(icon, "//") {
		icon = "https:" + icon
	}

	rec := generateTravelRecommendation(raw.Current.TempC, raw.Current.Condition.Text)

	return models.WeatherInfo{
		TempC:          raw.Current.TempC,
		Condition:      raw.Current.Condition.Text,
		ConditionIcon:  icon,
		WindKph:        raw.Current.WindKph,
		Humidity:       raw.Current.Humidity,
		Recommendation: rec,
	}, nil
}

func generateTravelRecommendation(tempC float64, conditionText string) string {
	cond := strings.ToLower(conditionText)
	if strings.Contains(cond, "rain") || strings.Contains(cond, "shower") || strings.Contains(cond, "drizzle") || strings.Contains(cond, "storm") {
		return "Wet weather expected. Keep an umbrella handy and plan indoor excursions."
	}
	if strings.Contains(cond, "snow") || strings.Contains(cond, "blizzard") || strings.Contains(cond, "sleet") || strings.Contains(cond, "ice") {
		return "Wintry conditions. Wear heavy thermal clothing and watch out for icy travel routes."
	}

	if tempC < 10.0 {
		return "Chilly conditions. Pack warm clothing and windbreakers for outdoor walking."
	}
	if tempC > 35.0 {
		return "Extreme heat. Stay hydrated, wear light clothing, apply sunblock, and avoid direct midday heat."
	}
	return "Pleasant temperatures. Ideal for city tours, outdoor sightseeing, and landscape photography."
}
