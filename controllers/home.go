package controllers

import (
	"TravelSphere/models"
)

// HomeController serves the landing page.
type HomeController struct {
	BaseController
}

// Get renders the home page with featured countries and popular attractions.
func (c *HomeController) Get() {
	c.Data["ActivePage"] = "home"
	c.TplName = "home.tpl"

	// Fetch some featured countries for the home display
	featuredSlugs := []string{"usa", "france", "japan", "bangladesh"}
	var featured []models.Country

	for _, slug := range featuredSlugs {
		country, err := CountrySvc.GetBySlug(slug)
		if err == nil {
			featured = append(featured, country)
		}
	}

	// If no countries are retrieved (e.g. cache not loaded yet), provide mock data to prevent empty UI
	if len(featured) == 0 {
		featured = []models.Country{
			{CommonName: "United States", Slug: "usa", Capital: "Washington D.C.", Region: "Americas", Flag: "🇺🇸", FormattedPopulation: "331,002,651"},
			{CommonName: "France", Slug: "france", Capital: "Paris", Region: "Europe", Flag: "🇫🇷", FormattedPopulation: "67,391,582"},
			{CommonName: "Japan", Slug: "japan", Capital: "Tokyo", Region: "Asia", Flag: "🇯🇵", FormattedPopulation: "125,836,021"},
			{CommonName: "Bangladesh", Slug: "bangladesh", Capital: "Dhaka", Region: "Asia", Flag: "🇧🇩", FormattedPopulation: "164,689,383"},
		}
	}

	c.Data["FeaturedCountries"] = featured

	// Popular attractions mock (representing top sights)
	c.Data["PopularAttractions"] = []models.Attraction{
		{Name: "Eiffel Tower", Kinds: "monuments,tourist_attractions", Dist: 0, Wikidata: "Q1234"},
		{Name: "Mount Fuji", Kinds: "natural,volcanoes", Dist: 0, Wikidata: "Q5678"},
		{Name: "Statue of Liberty", Kinds: "monuments,historical", Dist: 0, Wikidata: "Q9012"},
		{Name: "Lalbagh Fort", Kinds: "fortifications,historic", Dist: 0, Wikidata: "Q3456"},
	}
}
