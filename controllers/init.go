package controllers

import (
	"TravelSphere/services"
	"TravelSphere/utils"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

var (
	APIClient  *utils.APIClient
	CountrySvc *services.CountryService
)

func InitServices() {
	apiKey := beego.AppConfig.DefaultString("countries_api_key", "")
	APIClient = utils.NewAPIClient(10 * time.Second)
	CountrySvc = services.NewCountryService(APIClient, apiKey)
}
