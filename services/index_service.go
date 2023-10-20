package services

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/enum"
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

func GetNumberOfPages(totalPostsCount int64, numberOfPosts int) int {
	d := float64(totalPostsCount) / float64(numberOfPosts)
	return int(math.Ceil(d))
}

func GetPosts(c *fiber.Ctx, currentPage int, pageSize int, language string, tag string) []models.Post {

	var posts []models.Post
	dbQuery := initializers.DB.Order("created_at desc")

	isAuth := c.Locals("is_auth").(bool)
	showDrafts := getShowDrafts(c)

	if !isAuth {
		dbQuery.Where("is_draft = ?", "false")
	} else {
		dbQuery.Where("is_draft = ?", showDrafts)
	}

	if language != "All" {
		var languageKey uint8
		for k, v := range enum.LanguageEnumValues() {
			if v == language {
				languageKey = k
				break
			}
		}
		if languageKey != 0 {
			dbQuery.Where("language = ?", languageKey)
		}
	}

	dbQuery.Preload("Tags")
	dbQuery.Limit(pageSize).Offset(pageSize * (currentPage - 1))
	dbQuery.Find(&posts)

	if tag != "" { //i know... i know... this is way to ugly, but it is good enough for now
		var filteredPostsByTag []models.Post
		for _, post := range posts {
			for _, postTag := range post.Tags {
				if postTag.Name == tag {
					filteredPostsByTag = append(filteredPostsByTag, post)
				}
			}
		}
		return filteredPostsByTag
	}

	return posts
}

func GetTags(c *fiber.Ctx) []models.Tag {

	var tags []models.Tag
	dbQuery := initializers.DB.Order("id desc")

	dbQuery.Find(&tags)

	return tags
}

func GetPageSize(c *fiber.Ctx) int {
	queryParams := c.Queries()
	pageSize := queryParams["page_size"]

	if pageSize == "" {
		pageSize = "25"
	}

	pageSizeInt, _ := strconv.Atoi(pageSize)
	return pageSizeInt
}

func GetLanguage(c *fiber.Ctx) string {
	queryParams := c.Queries()
	language := queryParams["language"]

	if language == "" {
		language = "all"
	}

	return language
}

func GetTag(c *fiber.Ctx) string {
	queryParams := c.Queries()
	return queryParams["tag"]
}

func GetCurrentPage(c *fiber.Ctx) int {
	queryParams := c.Queries()
	page := queryParams["page"]

	if page == "" {
		page = "1"
	}

	pageInt, _ := strconv.Atoi(page)
	return pageInt
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
