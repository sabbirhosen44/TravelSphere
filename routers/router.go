package routers

import (
	"TravelSphere/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// SSR Page Routes
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/countries", &controllers.CountryController{})
	beego.Router("/countries/:slug", &controllers.CountryController{})
}
