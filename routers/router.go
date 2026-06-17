package routers

import (
	"TravelSphere/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// Register logging filter middleware
	beego.InsertFilter("*", beego.BeforeStatic, controllers.LogStartFilter)
	beego.InsertFilter("*", beego.FinishRouter, controllers.LogEndFilter)

	// Register authentication filter middleware
	beego.InsertFilter("/wishlist", beego.BeforeRouter, controllers.AuthFilter)
	beego.InsertFilter("/api/wishlist", beego.BeforeRouter, controllers.AuthFilter)
	beego.InsertFilter("/api/wishlist/*", beego.BeforeRouter, controllers.AuthFilter)
	beego.InsertFilter("/dashboard", beego.BeforeRouter, controllers.AuthFilter)

	// SSR Page Routes
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/countries", &controllers.CountryController{})
	beego.Router("/countries/:slug", &controllers.CountryController{})
	beego.Router("/dashboard", &controllers.DashboardController{})
	beego.Router("/wishlist", &controllers.WishlistController{})

	// Mock Authentication simulation routes
	beego.Router("/login", &controllers.AuthController{}, "get:Login")
	beego.Router("/logout", &controllers.AuthController{}, "get:Logout")
}
