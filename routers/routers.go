package routers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	fmt.Println(app)
}

// func helpcheck(r fiber.Router) {
// 	g := r.Group("")
// 	g.Get("/helpcheck", func(c *fiber.Ctx) error { return c.SendString("Hello World") })
// }
