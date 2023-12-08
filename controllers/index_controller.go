package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/dto"
	"github.com/marcoagpegoraro/marco_blog/models"
	"github.com/marcoagpegoraro/marco_blog/services"
)

var IndexController = IndexControllerStruct{}

type IndexControllerStruct struct {
}

func (controller IndexControllerStruct) Get(c *fiber.Ctx) error {
	currentPage := services.IndexService.GetCurrentPage(c)
	pageSize := services.IndexService.GetPageSize(c)
	language := services.IndexService.GetLanguage(c)
	tag := c.Queries()["tag"]
	showDrafts := services.IndexService.GetShowDrafts(c)
	isAuth := c.Locals("is_auth").(bool)

	channelPaginationButtons := make(chan []dto.PaginationButton)
	go services.IndexService.GetPaginationButtonsConcurrently(c, showDrafts, pageSize, currentPage, channelPaginationButtons)

	channelTags := make(chan []models.Tag)
	go services.IndexService.GetTagsConcurrently(c, isAuth, channelTags)

	channelPosts := make(chan []models.Post)
	go services.IndexService.GetPostsConcurrently(c, isAuth, currentPage, pageSize, language, tag, showDrafts, channelPosts)

	tags := <-channelTags
	posts := <-channelPosts
	paginationButtons := <-channelPaginationButtons

	if len(posts) == 0 && len(c.Queries()) != 0 {
		return c.RedirectToRoute("", fiber.Map{})
	}

	return c.Render("pages/index/index", fiber.Map{
		"title":              "Home",
		"posts":              posts,
		"tags":               tags,
		"base_url":           c.BaseURL(),
		"pagination_buttons": paginationButtons,
		"is_auth":            isAuth,
	}, "layouts/main")
}
