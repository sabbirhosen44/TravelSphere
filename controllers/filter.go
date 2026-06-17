package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
)

// LogStartFilter records the request start time.
func LogStartFilter(ctx *context.Context) {
	ctx.Input.SetData("startTime", time.Now())
}

// LogEndFilter calculates and logs request duration, HTTP method, and URL.
func LogEndFilter(ctx *context.Context) {
	startTimeVal := ctx.Input.GetData("startTime")
	if startTimeVal == nil {
		return
	}
	startTime, ok := startTimeVal.(time.Time)
	if !ok {
		return
	}

	duration := time.Since(startTime)
	method := ctx.Input.Method()
	uri := ctx.Input.URI()
	ip := ctx.Input.IP()
	status := ctx.ResponseWriter.Status

	logs.Info("[Logger] %s %s %s - Status: %d - Duration: %v", ip, method, uri, status, duration)
}

// AuthFilter protects access to wishlist and dashboard endpoints.
func AuthFilter(ctx *context.Context) {
	path := ctx.Input.URL()
	if path == "/" || strings.HasPrefix(path, "/static") || strings.HasPrefix(path, "/swagger") || path == "/countries" || strings.HasPrefix(path, "/countries/") {
		return
	}

	// We also have a simulation login/logout endpoint that we'll implement
	if path == "/login" || path == "/logout" {
		return
	}

	// Check forauthenticated=true cookie or a test header
	cookie, err := ctx.Request.Cookie("authenticated")
	isAuthenticated := false
	if err == nil && cookie.Value == "true" {
		isAuthenticated = true
	}

	// Allow query param or header override for testing/API requests
	if ctx.Input.Query("auth") == "1" || ctx.Input.Header("X-User-Authenticated") == "true" {
		isAuthenticated = true
	}

	if !isAuthenticated {
		if strings.HasPrefix(path, "/api/") {
			ctx.Output.SetStatus(http.StatusUnauthorized)
			_ = ctx.Output.JSON(map[string]interface{}{
				"error":   "unauthorized",
				"message": "Authentication required. Please set cookie 'authenticated=true' or pass header 'X-User-Authenticated: true'",
			}, false, false)
			return
		}

		ctx.Redirect(http.StatusFound, "/?error=unauthorized")
		return
	}
}
