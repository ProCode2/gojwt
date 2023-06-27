package routes

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")

	v1.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON("Hellow")
	})
}
