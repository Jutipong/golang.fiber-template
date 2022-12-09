package routers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"golang.fiber.template/model"
)

func PublicRoutes(app *fiber.App, db *model.ConnectDB) {
	fmt.Println(app)
}

// func helpcheck(r fiber.Router) {
// 	g := r.Group("")
// 	g.Get("/helpcheck", func(c *fiber.Ctx) error { return c.SendString("Hello World") })
// }
