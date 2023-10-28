package services

import (
	"fmt"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/dto"
	"github.com/marcoagpegoraro/marco_blog/helpers"
	"github.com/marcoagpegoraro/marco_blog/initializers"
	"github.com/marcoagpegoraro/marco_blog/models"
	"github.com/marcoagpegoraro/marco_blog/repositories"
	"github.com/patrickmn/go-cache"
)

var IndexService = IndexServiceStruct{}

type IndexServiceStruct struct {
}

func (service IndexServiceStruct) GetNumberOfPages(totalPostsCount int64, numberOfPosts int) int {
	d := float64(totalPostsCount) / float64(numberOfPosts)
	return int(math.Ceil(d))
}

func (service IndexServiceStruct) GetPostsConcurrently(c *fiber.Ctx, isAuth bool, currentPage int, pageSize int, language string, tag string, showDrafts bool, channelPosts chan []models.Post) {
	cacheKey := fmt.Sprintf("postsServiceGetPosts%t%d%d%s%s%t", isAuth, currentPage, pageSize, language, tag, showDrafts)

	var posts []models.Post
	if x, found := initializers.Cache.Get(cacheKey); found {
		posts = *x.(*[]models.Post)
	} else {
		posts = repositories.PostRepository.GetPosts(isAuth, currentPage, pageSize, language, tag, showDrafts)
		initializers.Cache.Set(cacheKey, &posts, cache.DefaultExpiration)
	}

	channelPosts <- posts
}

func (service IndexServiceStruct) GetTagsConcurrently(c *fiber.Ctx, isAuth bool, channelTags chan []models.Tag) {
	cacheKey := fmt.Sprintf("postsServiceGetTags%t", isAuth)

	var tags []models.Tag
	if x, found := initializers.Cache.Get(cacheKey); found {
		tags = *x.(*[]models.Tag)
	} else {
		tags = repositories.TagRepository.GetAllTags(isAuth)
		initializers.Cache.Set(cacheKey, &tags, cache.DefaultExpiration)
	}

	channelTags <- tags
}

func (service IndexServiceStruct) GetTotalPostsCount(c *fiber.Ctx, showDrafts bool) int64 {
	isAuth := c.Locals("is_auth").(bool)
	cacheKey := fmt.Sprintf("postsServiceGetTotalPostsCount%t%t", isAuth, showDrafts)

	var count int64
	if x, found := initializers.Cache.Get(cacheKey); found {
		count = *x.(*int64)
	} else {
		isAuth := c.Locals("is_auth").(bool)
		count = repositories.PostRepository.GetTotalPostsCount(isAuth, showDrafts)
		initializers.Cache.Set(cacheKey, &count, cache.DefaultExpiration)
	}

	return count
}

func (service IndexServiceStruct) GetPaginationButtonsConcurrently(c *fiber.Ctx, showDrafts bool, pageSize int, currentPage int, channelPaginationButtons chan []dto.PaginationButton) {
	totalPostsCount := service.GetTotalPostsCount(c, showDrafts)
	numberOfPages := service.GetNumberOfPages(totalPostsCount, pageSize)
	paginationButtons := helpers.CalculatePagination(numberOfPages, currentPage, 5)

	channelPaginationButtons <- paginationButtons
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
