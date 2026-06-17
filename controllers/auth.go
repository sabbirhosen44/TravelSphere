package controllers

import (
	"net/http"
	"time"
)


type AuthController struct {
	BaseController
}

// Login sets the authentication cookie and redirects to the dashboard.
func (c *AuthController) Login() {
	cookie := &http.Cookie{
		Name:     "authenticated",
		Value:    "true",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
	}
	http.SetCookie(c.Ctx.ResponseWriter, cookie)
	c.Redirect("/dashboard", http.StatusFound)
}

// Logout clears the authentication cookie and redirects to the home page.
func (c *AuthController) Logout() {
	cookie := &http.Cookie{
		Name:     "authenticated",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
	}
	http.SetCookie(c.Ctx.ResponseWriter, cookie)
	c.Redirect("/", http.StatusFound)
}
