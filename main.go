package main

import (
	"TravelSphere/controllers"
	_ "TravelSphere/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	// Initialize global singleton services
	controllers.InitServices()
	beego.Run()
}
