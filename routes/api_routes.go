package routes

import (
	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("it is on the air")
	})
}
