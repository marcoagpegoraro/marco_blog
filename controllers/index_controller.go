package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/helpers"
	"github.com/marcoagpegoraro/marco_blog/initializers"
	"github.com/marcoagpegoraro/marco_blog/models"
	"github.com/marcoagpegoraro/marco_blog/services"
	"github.com/patrickmn/go-cache"
)

var IndexController = IndexControllerStruct{}

type IndexControllerStruct struct {
}

func (controller IndexControllerStruct) Get(c *fiber.Ctx) error {
	currentPage := services.IndexService.GetCurrentPage(c)
	pageSize := services.IndexService.GetPageSize(c)
	language := services.IndexService.GetLanguage(c)
	tag := services.IndexService.GetTag(c)

	cacheKey := fmt.Sprintf("postsControllerGet%d%d%s%s", currentPage, pageSize, language, tag)
	var posts []models.Post
	if x, found := initializers.Cache.Get(cacheKey); found {
		posts = *x.(*[]models.Post)
	} else {
		posts = services.IndexService.GetPosts(c, currentPage, pageSize, language, tag)
		initializers.Cache.Set(cacheKey, &posts, cache.DefaultExpiration)
	}

	totalPostsCount := services.IndexService.GetTotalPostsCount(c)

	numberOfPages := services.IndexService.GetNumberOfPages(totalPostsCount, pageSize)
	paginationButtons := helpers.CalculatePagination(numberOfPages, currentPage, 5)

	tags := services.IndexService.GetTags(c)

	return c.Render("pages/index/index", fiber.Map{
		"title":              "Home",
		"posts":              posts,
		"tags":               tags,
		"base_url":           c.BaseURL(),
		"pagination_buttons": paginationButtons,
		"is_auth":            c.Locals("is_auth"),
	}, "layouts/main")
}
