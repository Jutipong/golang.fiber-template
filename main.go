package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
	"golang.fiber.template/routers"
)

func main() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "*",
		AllowHeaders:     "*",
		AllowCredentials: true}))
	app.Use(recover.New())
	app.Use(requestid.New(requestid.Config{
		Generator: func() string { return uuid.New().String() },
		// ContextKey: enums.TransactionId,
	}))

	// routers.PublicRoutes(app, db)
	routers.NotFoundRoute(app)
}
