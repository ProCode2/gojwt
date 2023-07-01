package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/procode2/etir/handlers"
)

func RegisterRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")

	// authentication routes
	auth := v1.Group("/auth")

	auth.Get("/", handlers.HanldeGetAuthUser)
	auth.Post("/", handlers.HanldeCreateAuthUser)
	auth.Delete("/", handlers.HandleDeleteAuthUser)
}
