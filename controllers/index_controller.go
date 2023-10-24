package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/helpers"
	"github.com/marcoagpegoraro/marco_blog/initializers"
	"github.com/marcoagpegoraro/marco_blog/services"
	"github.com/patrickmn/go-cache"
)

var IndexController = IndexControllerStruct{}

type IndexControllerStruct struct {
}

func (controller IndexControllerStruct) Get(c *fiber.Ctx) error {
	currentPage := services.GetCurrentPage(c)
	pageSize := services.GetPageSize(c)
	language := services.GetLanguage(c)
	tag := services.GetTag(c)

	cacheKeyPosts := fmt.Sprintf("postsControllerGet%d%d%s%s", currentPage, pageSize, language, tag)
	posts, found := initializers.Cache.Get(cacheKeyPosts)

	fmt.Println(found)
	if !found {
		posts = services.GetPosts(c, currentPage, pageSize, language, tag)
		initializers.Cache.Set(cacheKeyPosts, posts, cache.DefaultExpiration)
	}

	totalPostsCount := services.GetTotalPostsCount(c)

	numberOfPages := services.GetNumberOfPages(totalPostsCount, pageSize)
	paginationButtons := helpers.CalculatePagination(numberOfPages, currentPage, 5)

	tags := services.GetTags(c)

	return c.Render("pages/index/index", fiber.Map{
		"title":              "Home",
		"posts":              posts,
		"tags":               tags,
		"base_url":           c.BaseURL(),
		"pagination_buttons": paginationButtons,
		"is_auth":            c.Locals("is_auth"),
	}, "layouts/main")
}
