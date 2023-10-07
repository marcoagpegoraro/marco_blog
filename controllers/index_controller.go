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
		pageSize = "25"
	}

	pageSizeInt, _ := strconv.Atoi(pageSize)

	var posts []models.Post
	initializers.DB.Order("created_at desc").Limit(pageSizeInt).Where("is_draft = ?", "false").Preload("Tags").Find(&posts)

	return c.Render("pages/index/index", fiber.Map{
		"title": "Home",
		"posts": posts,
	}, "layouts/main")
}
