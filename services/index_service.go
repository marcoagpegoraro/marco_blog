package services

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/enum"
	"github.com/marcoagpegoraro/marco_blog/initializers"
	"github.com/marcoagpegoraro/marco_blog/models"
)

var IndexService = IndexServiceStruct{}

type IndexServiceStruct struct {
}

func (service IndexServiceStruct) GetTotalPostsCount(c *fiber.Ctx, showDrafts bool) int64 {
	isAuth := c.Locals("is_auth").(bool)

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

func (service IndexServiceStruct) GetNumberOfPages(totalPostsCount int64, numberOfPosts int) int {
	d := float64(totalPostsCount) / float64(numberOfPosts)
	return int(math.Ceil(d))
}

func (service IndexServiceStruct) GetPosts(c *fiber.Ctx, currentPage int, pageSize int, language string, tag string, showDrafts bool) []models.Post {

	var posts []models.Post
	dbQuery := initializers.DB.Preload("Tags")

	if isAuth := c.Locals("is_auth").(bool); isAuth {
		dbQuery.Where("is_draft = ?", showDrafts)
		if showDrafts {
			dbQuery.Order("created_at desc")
		} else {
			dbQuery.Order("publicated_at desc")
		}
	} else {
		dbQuery.Where("is_draft = ?", "false")
		dbQuery.Order("publicated_at desc")
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

	dbQuery.Limit(pageSize).Offset(pageSize * (currentPage - 1))
	dbQuery.Find(&posts)

	if tag != "" { //i know... i know... this is way to ugly, but it is good enough for now
		var filteredPostsByTag []models.Post
		for _, post := range posts {
			for _, postTag := range post.Tags {
				if postTag.Name == tag {
					filteredPostsByTag = append(filteredPostsByTag, post)
					break
				}
			}
		}
		return filteredPostsByTag
	}

	return posts
}

func (service IndexServiceStruct) GetTags(c *fiber.Ctx) []models.Tag {

	sqlQuery := `
    SELECT distinct t.name
    FROM tags t
    LEFT OUTER JOIN posts_tags pt
    ON t.name = pt.tag_name
    LEFT OUTER JOIN posts p
    ON pt.post_id  = p.id  
    `

	var tags []models.Tag

	if isAuth := c.Locals("is_auth").(bool); !isAuth {
		sqlQuery += ` WHERE  p.is_draft = false`
	}

	sqlQuery += ` ORDER BY t.name`

	initializers.DB.Raw(sqlQuery).Scan(&tags)

	return tags
}

func (service IndexServiceStruct) GetTagsConcurrently(c *fiber.Ctx, channelTags chan []models.Tag) {
	tags := service.GetTags(c)
	channelTags <- tags
}

func (service IndexServiceStruct) GetPageSize(c *fiber.Ctx) int {
	queryParams := c.Queries()
	pageSize := queryParams["page_size"]

	if pageSize == "" {
		pageSize = "25"
	}

	pageSizeInt, _ := strconv.Atoi(pageSize)
	return pageSizeInt
}

func (service IndexServiceStruct) GetLanguage(c *fiber.Ctx) string {
	queryParams := c.Queries()
	language := queryParams["language"]

	if language == "" {
		language = "all"
	}

	return language
}

func (service IndexServiceStruct) GetTag(c *fiber.Ctx) string {
	queryParams := c.Queries()
	return queryParams["tag"]
}

func (service IndexServiceStruct) GetCurrentPage(c *fiber.Ctx) int {
	queryParams := c.Queries()
	page := queryParams["page"]

	if page == "" {
		page = "1"
	}

	pageInt, _ := strconv.Atoi(page)
	return pageInt
}

func (service IndexServiceStruct) GetShowDrafts(c *fiber.Ctx) bool {
	queryParams := c.Queries()
	showDrafts := queryParams["show_drafts"]

	showDraftsBool, err := strconv.ParseBool(showDrafts)
	if err != nil {
		showDraftsBool = false
	}

	return showDraftsBool
}
