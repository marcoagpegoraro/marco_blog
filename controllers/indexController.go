package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/initializers"
	"github.com/marcoagpegoraro/marco_blog/models"
)

func GetIndex(c *fiber.Ctx) error {

	queryParams := c.Queries()
	pageSize := queryParams["page_size"]

	if pageSize == "" {
		pageSize = "10"
	}

	pageSizeInt, _ := strconv.Atoi(pageSize)

	var posts []models.Post
	initializers.DB.Order("created_at desc").Limit(pageSizeInt).Find(&posts)

	return c.Render("index/index", fiber.Map{
		"Title": "Home",
		"Posts": posts,
	}, "layouts/main")
}
