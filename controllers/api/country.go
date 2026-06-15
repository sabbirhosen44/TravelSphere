package api

import (
	"TravelSphere/controllers"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

type CountryController struct {
	beego.Controller
}

func (c *CountryController) Get() {
	slug := c.Ctx.Input.Param(":slug")
	if slug != "" {
		c.getDetails(slug)
		return
	}

	search := c.GetString("search")
	region := c.GetString("region")

	list, err := controllers.CountrySvc.List(search, region)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": err.Error()}
		_ = c.ServeJSON()
		return
	}

	c.Data["json"] = list
	_ = c.ServeJSON()
}

func (c *CountryController) getDetails(slug string) {
	country, err := controllers.CountrySvc.GetBySlug(slug)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusNotFound)
		c.Data["json"] = map[string]string{"error": "country not found"}
		_ = c.ServeJSON()
		return
	}

	c.Data["json"] = country
	_ = c.ServeJSON()
}
