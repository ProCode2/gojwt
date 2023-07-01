package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/procode2/etir/routes"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routes.RegisterRoutes(app)

	app.Listen(":3000")
}
