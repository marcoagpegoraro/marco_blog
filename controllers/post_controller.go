package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/dto"
	"github.com/marcoagpegoraro/marco_blog/enum"
	"github.com/marcoagpegoraro/marco_blog/initializers"
	"github.com/marcoagpegoraro/marco_blog/repositories"
	"github.com/marcoagpegoraro/marco_blog/services"
)

var PostController = PostControllerStruct{}

type PostControllerStruct struct {
}

func (controller PostControllerStruct) Get(c *fiber.Ctx) error {
	pathPrefix := c.Locals("layout_path_prefix").(string)

	return c.Render(pathPrefix+"pages/posts/index", fiber.Map{
		"title":     "Create post",
		"languages": enum.LanguageEnumValues(),
		"is_auth":   c.Locals("is_auth"),
	}, pathPrefix+"layouts/main")
}

func (controller PostControllerStruct) GetOne(c *fiber.Ctx) error {
	id, err := services.PostService.GetIdParamFromUrl(c)
	if err != nil {
		return c.RedirectToRoute("", fiber.Map{
			"messages": []dto.MessageDto{{Message: "Post not found", Type: "warning"}},
		})
	}

	post := services.PostService.GetPostById(id)

	if isAuth := c.Locals("is_auth").(bool); !isAuth && post.IsDraft {
		c.RedirectToRoute("", fiber.Map{})
	}
	pathPrefix := c.Locals("layout_path_prefix").(string)

	return c.Render(pathPrefix+"pages/posts/one", fiber.Map{
		"title":        post.Title,
		"post":         post,
		"base_url":     c.BaseURL(),
		"original_url": c.OriginalURL(),
		"is_auth":      c.Locals("is_auth"),
	}, pathPrefix+"layouts/main")
}

func (controller PostControllerStruct) Post(c *fiber.Ctx) error {
	postModel, err := services.PostService.HandlePostRequestPost(c)
	if err != nil {
		return err
	}

	repositories.PostRepository.Insert(&postModel)

	initializers.Cache.Flush()

	return c.Redirect(fmt.Sprintf("/posts/%d", postModel.Id))
}

func (controller PostControllerStruct) GetEditPost(c *fiber.Ctx) error {
	id, err := services.PostService.GetIdParamFromUrl(c)
	if err != nil {
		return c.RedirectToRoute("", fiber.Map{
			"messages": []dto.MessageDto{{Message: "Post not found", Type: "warning"}},
		})
	}

	post := repositories.PostRepository.GetPostById(id)

	pathPrefix := c.Locals("layout_path_prefix").(string)

	return c.Render(pathPrefix+"pages/posts/index", fiber.Map{
		"title":     "Edit post",
		"post":      post,
		"languages": enum.LanguageEnumValues(),
		"is_auth":   c.Locals("is_auth"),
	}, pathPrefix+"layouts/main")
}

func (controller PostControllerStruct) PostEditPost(c *fiber.Ctx) error {
	postModel, err := services.PostService.HandlePostRequestPost(c)
	if err != nil {
		return err
	}

	services.PostService.UpdatePost(postModel)

	initializers.Cache.Flush()

	return c.Redirect(fmt.Sprintf("/posts/%d", postModel.Id))
}
