package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

// BaseController provides template layout binding and common Prepare logic.
type BaseController struct {
	web.Controller
}

// Prepare binds the layout, partials, and authentication state for SSR routes.
func (c *BaseController) Prepare() {
	// Config layouts
	c.Layout = "layout.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "header.tpl"
	c.LayoutSections["Footer"] = "footer.tpl"

	// Determine authentication state for UI displays
	cookie, err := c.Ctx.Request.Cookie("authenticated")
	isAuthenticated := false
	if err == nil && cookie.Value == "true" {
		isAuthenticated = true
	}
	c.Data["IsAuthenticated"] = isAuthenticated

	// Default active page highlighting
	c.Data["ActivePage"] = ""
	c.Data["Year"] = 2026
	c.Data["ErrorMessage"] = c.GetString("error")
	c.Data["SuccessMessage"] = c.GetString("success")
}
