package models

type Country struct {
	Name                string `json:"name"`
	CommonName          string `json:"common_name"`
	Slug                string `json:"slug"`
	Flag                string `json:"flag"`
	Capital             string `json:"capital"`
	Region              string `json:"region"`
	Population          int64  `json:"population"`
	FormattedPopulation string `json:"formatted_population"`
	Currencies          string `json:"currencies"`
	Languages           string `json:"languages"`
	LatLng              []float64 `json:"latlng"`
}

// Attraction represents a tourist attraction from OpenTripMap.
type Attraction struct {
	Name     string  `json:"name"`
	Kinds    string  `json:"kinds"`
	Dist     float64 `json:"dist"`
	Wikidata string  `json:"wikidata"`
}

// WeatherInfo represents live weather conditions from WeatherAPI.
type WeatherInfo struct {
	TempC          float64 `json:"temp_c"`
	Condition      string  `json:"condition"`
	ConditionIcon  string  `json:"condition_icon"`
	WindKph        float64 `json:"wind_kph"`
	Humidity       int     `json:"humidity"`
	Recommendation string  `json:"recommendation"`
}
