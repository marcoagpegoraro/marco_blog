package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func OnlyAuthenticated(c *fiber.Ctx) error {
	isAuth := c.Locals("is_auth").(bool)

	if isAuth {
		return c.Next()
	} else {
		return c.Redirect(c.BaseURL())
	}
}
