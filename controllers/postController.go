package controllers

import "github.com/gofiber/fiber/v2"

func PostIndex(c *fiber.Ctx) error {
	return c.Render("posts/index", fiber.Map{
		"Title": "Hello, World!",
	}, "layouts/main")
}
