package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/services"
)

func GetIndex(c *fiber.Ctx) error {
	currentPage := services.GetCurrentPage(c)

	posts := services.GetPosts(c, currentPage)
	totalPostsCount := services.GetTotalPostsCount(c)

	numberOfPages := services.GetNumberOfPages(totalPostsCount, len(posts))

	return c.Render("pages/index/index", fiber.Map{
		"title":         "Home",
		"posts":         posts,
		"currentPage":   currentPage,
		"numberOfPages": numberOfPages,
		"is_auth":       c.Locals("is_auth"),
	}, "layouts/main")
}
