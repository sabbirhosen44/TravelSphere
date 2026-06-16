package controllers

import (
	"fmt"
	"math"
	"net/http"
	"strings"

	"TravelSphere/models"

	"github.com/beego/beego/v2/server/web"
)

type CountryController struct {
	BaseController
}

func (c *CountryController) Get() {
	slug := c.Ctx.Input.Param(":slug")

	if slug == "" {
		c.renderExplorer()
	} else {
		c.renderDetails(slug)
	}
}

func (c *CountryController) renderExplorer() {
	c.Data["ActivePage"] = "countries"
	c.TplName = "countries.tpl"

	countriesBaseURL, _ := web.AppConfig.String("countries_base_url")
	_ = countriesBaseURL

	searchQuery := strings.TrimSpace(c.GetString("search"))
	regionQuery := strings.TrimSpace(c.GetString("region"))

	page, err := c.GetInt("page")
	if err != nil || page < 1 {
		page = 1
	}
	pageSize := 12

	list, err := CountrySvc.List(searchQuery, regionQuery)
	if err != nil {
		if c.Ctx.Input.URL() == "/api/countries" {
			c.CustomAbort(http.StatusInternalServerError, "Failed to load backend metrics.")
			return
		}
		c.Data["Countries"] = []models.Country{}
		c.Data["Error"] = "Failed to load countries from external API."
		return
	}

	totalItems := len(list)
	totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))
	if totalPages == 0 {
		totalPages = 1
	}

	if page > totalPages {
		page = totalPages
	}

	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize
	if endIndex > totalItems {
		endIndex = totalItems
	}

	var paginatedList []models.Country
	if startIndex < totalItems {
		paginatedList = list[startIndex:endIndex]
	} else {
		paginatedList = []models.Country{}
	}

	hasPrev := page > 1
	hasNext := page < totalPages
	prevPage := page - 1
	nextPage := page + 1

	if c.Ctx.Input.URL() == "/api/countries" {
		c.Data["json"] = map[string]interface{}{
			"countries":      paginatedList,
			"total_pages":    totalPages,
			"current_page":   page,
			"has_prev":       hasPrev,
			"has_next":       hasNext,
			"prev_page":      prevPage,
			"next_page":      nextPage,
			"current_search": searchQuery,
			"current_region": regionQuery,
		}
		c.ServeJSON()
		return
	}

	c.Data["Countries"] = paginatedList
	c.Data["TotalPages"] = totalPages
	c.Data["CurrentPage"] = page
	c.Data["HasPrev"] = hasPrev
	c.Data["HasNext"] = hasNext
	c.Data["PrevPage"] = prevPage
	c.Data["NextPage"] = nextPage
	c.Data["CurrentSearch"] = searchQuery
	c.Data["CurrentRegion"] = regionQuery
}

func (c *CountryController) renderDetails(slug string) {
	c.Data["ActivePage"] = "countries"

	tripBaseURL, _ := web.AppConfig.String("trip_base_url")
	_ = tripBaseURL

	country, err := CountrySvc.GetBySlug(slug)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
		c.Data["Slug"] = slug
		c.TplName = "404.tpl"
		return
	}

	c.TplName = "destination.tpl"
	c.Data["Country"] = country

	weather, err := WeatherSvc.GetWeather(country.Capital)
	if err != nil {
		weather = models.WeatherInfo{
			TempC:          20.0,
			Condition:      "Weather service temporarily unavailable",
			ConditionIcon:  "https://cdn.weatherapi.com/weather/64x64/day/116.png",
			Recommendation: "Pack weather-appropriate clothing for local travel.",
		}
	}
	c.Data["Weather"] = weather

	lat := 0.0
	lon := 0.0

	fmt.Println(country)

	if len(country.LatLng) >= 2 {
		lat = country.LatLng[0]
		lon = country.LatLng[1]
	}

	attractions, err := CountrySvc.GetAttractions(lat, lon)
	if err != nil {
		attractions = []models.Attraction{
			{Name: "Attractions information currently offline", Kinds: "offline"},
		}
	}

	if len(attractions) > 6 {
		attractions = attractions[:6]
	}
	c.Data["Attractions"] = attractions

	wishlistItems, err := WishlistSvc.List()
	isAdded := false
	wishlistItemID := ""
	if err == nil {
		for _, item := range wishlistItems {
			if strings.EqualFold(item.CountryName, country.CommonName) {
				isAdded = true
				wishlistItemID = item.ID
				break
			}
		}
	}
	c.Data["IsAddedToWishlist"] = isAdded
	c.Data["WishlistItemID"] = wishlistItemID
}
