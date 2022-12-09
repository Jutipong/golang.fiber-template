package main

import (
	"golang.fiber.template/pkg/config"
	"golang.fiber.template/pkg/server"
	"golang.fiber.template/routers"
)

func main() {
	cf := config.Init()
	app := App()

	// routers.PublicRoutes(app, db)
	routers.NotFoundRoute(app)
	server.Start(app, cf)
}
