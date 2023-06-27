package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/procode2/etir/routes"
)

func main() {
	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World")
	// })

	app.Static("/", "./public")

	routes.RegisterRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
