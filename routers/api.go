package routers

import (
	"TravelSphere/controllers/api"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/api/countries", &api.CountryController{})
	beego.Router("/api/countries/:slug", &api.CountryController{})
	beego.Router("/api/wishlist", &api.WishlistController{})
	beego.Router("/api/wishlist/:id", &api.WishlistController{})
}
