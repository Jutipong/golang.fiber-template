package routers

import "github.com/gofiber/fiber/v2"

func NotFoundRoute(app *fiber.App) {
	app.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "sorry, endpoint is not found",
		})
	})
}
