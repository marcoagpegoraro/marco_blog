package controllers

import "github.com/gofiber/fiber/v2"

func GetIndex(c *fiber.Ctx) error {
	return c.Render("index/index", fiber.Map{
		"Title": "Home",
	}, "layouts/main")
}
