package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/services"
)

func GetIndex(c *fiber.Ctx) error {
	posts := services.GetPosts(c)
	totalPostsCount := services.GetTotalPostsCount(c)

	return c.Render("pages/index/index", fiber.Map{
		"title":           "Home",
		"posts":           posts,
		"totalPostsCount": totalPostsCount,
		"is_auth":         c.Locals("is_auth"),
	}, "layouts/main")
}
