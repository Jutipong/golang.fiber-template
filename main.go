package main

import (
	"golang.fiber.template/model"
	"golang.fiber.template/pkg/config"
	"golang.fiber.template/pkg/server"
	"golang.fiber.template/routers"
)

func main() {
	cf := config.Init()
	db := new(model.ConnectDB)
	defer func() {}()

	app := server.App()
	routers.NotFoundRoute(app)
	routers.PublicRoutes(app, db)
	server.Start(app, cf)
}
