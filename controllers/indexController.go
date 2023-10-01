package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/initializers"
	"github.com/marcoagpegoraro/marco_blog/models"
)

func GetIndex(c *fiber.Ctx) error {

	var posts []models.Post
	initializers.DB.Find(&posts)

	return c.Render("index/index", fiber.Map{
		"Title": "Home",
		"Posts": posts,
	}, "layouts/main")
}
