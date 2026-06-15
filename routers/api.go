package routers

import (
	"TravelSphere/controllers/api"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/api/countries", &api.CountryController{})
}
