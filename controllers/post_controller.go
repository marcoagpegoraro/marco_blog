package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/enum"
	"github.com/marcoagpegoraro/marco_blog/helpers"
	"github.com/marcoagpegoraro/marco_blog/initializers"
	"github.com/marcoagpegoraro/marco_blog/models"
)

func GetPostIndex(c *fiber.Ctx) error {
	return c.Render("pages/posts/index", fiber.Map{
		"title":     "Create new post",
		"languages": enum.LanguageEnumValues(),
		"is_auth":   c.Locals("is_auth"),
	}, "layouts/main")
}

func GetOnePostIndex(c *fiber.Ctx) error {
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
		"title":   "Create new post",
		"post":    post,
		"is_auth": c.Locals("is_auth"),
	}, "layouts/main")
}

func PostPostIndex(c *fiber.Ctx) error {
	postModel, err := helpers.HandlePostRequestPost(c)
	if err != nil {
		return err
	}

	initializers.DB.Create(&postModel)

	return c.Redirect(fmt.Sprintf("/posts/%d", postModel.Id))
}

func GetPostUpdate(c *fiber.Ctx) error {
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

func PostPostUpdate(c *fiber.Ctx) error {
	postModel, err := helpers.HandlePostRequestPost(c)
	if err != nil {
		return err
	}

	saveResult := initializers.DB.Omit("created_at").Save(&postModel)

	fmt.Println(saveResult)

	return c.Redirect(fmt.Sprintf("/posts/%d", postModel.Id))
}
