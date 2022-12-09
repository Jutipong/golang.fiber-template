package server

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "*",
		AllowHeaders:     "*",
		AllowCredentials: true}))
	// app.Use(requestid.New(requestid.Config{ContextKey: "TransactionId"}))
	return app
}
