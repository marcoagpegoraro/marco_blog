package services

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/initializers"
	"github.com/marcoagpegoraro/marco_blog/models"
)

func GetTotalPostsCount(c *fiber.Ctx) int64 {
	isAuth := c.Locals("is_auth").(bool)
	showDrafts := getShowDrafts(c)

	var count int64
	dbQuery := initializers.DB.Model(&models.Post{})

	if !isAuth {
		dbQuery.Where("is_draft = ?", "false")
	} else {
		dbQuery.Where("is_draft = ?", showDrafts)
	}

	dbQuery.Count(&count)

	return count
}

func GetPosts(c *fiber.Ctx) []models.Post {

	pageSize := getPageSize(c)

	var posts []models.Post
	dbQuery := initializers.DB.Order("created_at desc")

	isAuth := c.Locals("is_auth").(bool)
	showDrafts := getShowDrafts(c)

	if !isAuth {
		dbQuery.Where("is_draft = ?", "false")
	} else {
		dbQuery.Where("is_draft = ?", showDrafts)
	}

	dbQuery.Limit(pageSize)
	dbQuery.Preload("Tags")
	dbQuery.Find(&posts)

	return posts
}

func getPageSize(c *fiber.Ctx) int {
	queryParams := c.Queries()
	pageSize := queryParams["page_size"]

	if pageSize == "" {
		pageSize = "25"
	}

	pageSizeInt, _ := strconv.Atoi(pageSize)
	return pageSizeInt
}

func getShowDrafts(c *fiber.Ctx) bool {
	queryParams := c.Queries()
	showDrafts := queryParams["show_drafts"]

	showDraftsBool, err := strconv.ParseBool(showDrafts)
	if err != nil {
		showDraftsBool = false
	}

	return showDraftsBool
}
