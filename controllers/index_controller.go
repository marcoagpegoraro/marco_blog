package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/helpers"
	"github.com/marcoagpegoraro/marco_blog/services"
)

func GetIndex(c *fiber.Ctx) error {
	currentPage := services.GetCurrentPage(c)
	pageSize := services.GetPageSize(c)
	language := services.GetLanguage(c)
	tag := services.GetTag(c)

	posts := services.GetPosts(c, currentPage, pageSize, language, tag)
	totalPostsCount := services.GetTotalPostsCount(c)

	numberOfPages := services.GetNumberOfPages(totalPostsCount, pageSize)
	paginationButtons := helpers.CalculatePagination(numberOfPages, currentPage, 5)

	tags := services.GetTags(c)

	return c.Render("pages/index/index", fiber.Map{
		"title":              "Home",
		"posts":              posts,
		"tags":               tags,
		"pagination_buttons": paginationButtons,
		"is_auth":            c.Locals("is_auth"),
	}, "layouts/main")
}
