package controllers

import "github.com/gofiber/fiber/v2"

func PostIndex(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}
