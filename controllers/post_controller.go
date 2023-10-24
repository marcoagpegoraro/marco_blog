package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/enum"
	"github.com/marcoagpegoraro/marco_blog/helpers"
	"github.com/marcoagpegoraro/marco_blog/initializers"
	"github.com/marcoagpegoraro/marco_blog/models"
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
	id, err := helpers.GetIdParamFromUrl(c)
	if err != nil {
		return err
	}

	var post models.Post
	initializers.DB.Where("id = ?", id).Preload("Tags").First(&post)

	isAuth := c.Locals("is_auth").(bool)

	if !isAuth && post.IsDraft {
		c.RedirectToRoute("", fiber.Map{})
	}

	return c.Render("pages/posts/one", fiber.Map{
		"title":   post.Title,
		"post":    post,
		"is_auth": c.Locals("is_auth"),
	}, "layouts/main")
}

func (controller PostControllerStruct) Post(c *fiber.Ctx) error {
	postModel, err := helpers.HandlePostRequestPost(c)
	if err != nil {
		return err
	}

	initializers.DB.Create(&postModel)

	return c.Redirect(fmt.Sprintf("/posts/%d", postModel.Id))
}

func (controller PostControllerStruct) GetEditPost(c *fiber.Ctx) error {
	id, err := helpers.GetIdParamFromUrl(c)
	if err != nil {
		return err
	}

	var post models.Post
	initializers.DB.Where("id = ?", id).Preload("Tags").First(&post)

	return c.Render("pages/posts/index", fiber.Map{
		"title":     "Create new post",
		"post":      post,
		"languages": enum.LanguageEnumValues(),
		"is_auth":   c.Locals("is_auth"),
	}, "layouts/main")
}

func (controller PostControllerStruct) PostEditPost(c *fiber.Ctx) error {
	postModel, err := helpers.HandlePostRequestPost(c)
	if err != nil {
		return err
	}

	tagsToBeDeleted := helpers.GetTagsToBeDeleted(postModel.Id, postModel)

	initializers.DB.Model(&postModel).Association("Tags").Delete(tagsToBeDeleted)
	initializers.DB.Omit("created_at").Save(&postModel)

	return c.Redirect(fmt.Sprintf("/posts/%d", postModel.Id))
}
