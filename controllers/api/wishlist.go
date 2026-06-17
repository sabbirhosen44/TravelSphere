package api

import (
	"encoding/json"
	"net/http"

	"TravelSphere/controllers"
	"TravelSphere/utils"

	"github.com/beego/beego/v2/server/web"
)

// WishlistController handles CRUD actions on user wishlists.
type WishlistController struct {
	web.Controller
}

// Prepare ensures request bodies are copied for JSON parsing.
func (c *WishlistController) Prepare() {
	// Enable request body copying for JSON parsing
	c.Ctx.Input.CopyBody(1 << 20) // 1MB limit
}

// WishlistPayload defines the expected JSON input schema.
type WishlistPayload struct {
	CountryName string `json:"country_name"`
	Note        string `json:"note"`
	Status      string `json:"status"`
}


func (c *WishlistController) Get() {
	list, err := controllers.WishlistSvc.List()
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": err.Error()}
		_ = c.ServeJSON()
		return
	}

	c.Data["json"] = list
	_ = c.ServeJSON()
}

func (c *WishlistController) Post() {
	var payload WishlistPayload
	
	// Parse body (JSON or form fallback)
	if len(c.Ctx.Input.RequestBody) > 0 {
		_ = json.Unmarshal(c.Ctx.Input.RequestBody, &payload)
	}
	if payload.CountryName == "" {
		payload.CountryName = c.GetString("country_name")
		payload.Note = c.GetString("note")
		payload.Status = c.GetString("status")
	}

	// Validate input
	if err := utils.ValidateWishlistInput(payload.CountryName, payload.Status, payload.Note); err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": err.Error()}
		_ = c.ServeJSON()
		return
	}

	item, err := controllers.WishlistSvc.Create(payload.CountryName, payload.Note, payload.Status)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": err.Error()}
		_ = c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusCreated)
	c.Data["json"] = item
	_ = c.ServeJSON()
}


func (c *WishlistController) Put() {
	id := c.Ctx.Input.Param(":id")
	if id == "" {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "missing item ID"}
		_ = c.ServeJSON()
		return
	}

	var payload WishlistPayload
	if len(c.Ctx.Input.RequestBody) > 0 {
		_ = json.Unmarshal(c.Ctx.Input.RequestBody, &payload)
	}
	if payload.Status == "" {
		payload.Status = c.GetString("status")
		payload.Note = c.GetString("note")
	}

	// Validate updated input (CountryName is bypassed for updates, set dummy valid name for validation)
	if err := utils.ValidateWishlistInput("France", payload.Status, payload.Note); err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": err.Error()}
		_ = c.ServeJSON()
		return
	}

	item, err := controllers.WishlistSvc.Update(id, payload.Note, payload.Status)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusNotFound)
		c.Data["json"] = map[string]string{"error": err.Error()}
		_ = c.ServeJSON()
		return
	}

	c.Data["json"] = item
	_ = c.ServeJSON()
}


func (c *WishlistController) Delete() {
	id := c.Ctx.Input.Param(":id")
	if id == "" {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "missing item ID"}
		_ = c.ServeJSON()
		return
	}

	err := controllers.WishlistSvc.Delete(id)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusNotFound)
		c.Data["json"] = map[string]string{"error": err.Error()}
		_ = c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]string{
		"message": "success",
		"id":      id,
	}
	_ = c.ServeJSON()
}
