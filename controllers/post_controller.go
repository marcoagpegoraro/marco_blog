package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/enum"
	"github.com/marcoagpegoraro/marco_blog/initializers"
	"github.com/marcoagpegoraro/marco_blog/repositories"
	"github.com/marcoagpegoraro/marco_blog/services"
)

var PostController = PostControllerStruct{}

type PostControllerStruct struct {
}

func (controller PostControllerStruct) Get(c *fiber.Ctx) error {
	return c.Render("pages/posts/index", fiber.Map{
		"title":     "Create new post",
		"languages": enum.LanguageEnumValues(),
		"is_auth":   c.Locals("is_auth"),
	}, "layouts/main")
}

func (controller PostControllerStruct) GetOne(c *fiber.Ctx) error {
	id, err := services.PostService.GetIdParamFromUrl(c)
	if err != nil {
		return err
	}

	post := services.PostService.GetPostById(id)

	if isAuth := c.Locals("is_auth").(bool); !isAuth && post.IsDraft {
		c.RedirectToRoute("", fiber.Map{})
	}

	return c.Render("pages/posts/one", fiber.Map{
		"title":   post.Title,
		"post":    post,
		"is_auth": c.Locals("is_auth"),
	}, "layouts/main")
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
		return err
	}

	post := repositories.PostRepository.GetPostById(id)

	return c.Render("pages/posts/index", fiber.Map{
		"title":     "Create new post",
		"post":      post,
		"languages": enum.LanguageEnumValues(),
		"is_auth":   c.Locals("is_auth"),
	}, "layouts/main")
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
