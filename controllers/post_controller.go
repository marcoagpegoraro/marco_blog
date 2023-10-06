package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/enum"
	"github.com/marcoagpegoraro/marco_blog/helpers"
	"github.com/marcoagpegoraro/marco_blog/initializers"
	"github.com/marcoagpegoraro/marco_blog/models"
)

func GetPostIndex(c *fiber.Ctx) error {
	return c.Render("posts/index", fiber.Map{
		"title":     "Create new post",
		"languages": enum.LanguageEnumValues(),
	}, "layouts/main")
}

func GetOnePostIndex(c *fiber.Ctx) error {
	id, err := helpers.GetIdParamFromUrl(c)
	if err != nil {
		return err
	}

	var post models.Post
	initializers.DB.Where("id = ?", id).Preload("Tags").First(&post)

	return c.Render("posts/one", fiber.Map{
		"title": "Create new post",
		"post":  post,
	}, "layouts/main")
}

func PostPostIndex(c *fiber.Ctx) error {
	postModel, err := helpers.HandlePostRequestPost(c)
	if err != nil {
		return err
	}

	initializers.DB.Create(&postModel)

	return c.Render("posts/index", fiber.Map{
		"title":     "Create new post",
		"languages": enum.LanguageEnumValues(),
	}, "layouts/main")
}

func GetPostUpdate(c *fiber.Ctx) error {
	id, err := helpers.GetIdParamFromUrl(c)
	if err != nil {
		return err
	}

	var post models.Post
	initializers.DB.Where("id = ?", id).Preload("Tags").First(&post)

	return c.Render("posts/index", fiber.Map{
		"title":     "Create new post",
		"post":      post,
		"languages": enum.LanguageEnumValues(),
	}, "layouts/main")
}

func PostPostUpdate(c *fiber.Ctx) error {
	postModel, err := helpers.HandlePostRequestPost(c)
	if err != nil {
		return err
	}

	initializers.DB.Omit("created_at").Save(&postModel)

	return c.Render("posts/index", fiber.Map{
		"title":     "Create new post",
		"languages": enum.LanguageEnumValues(),
	}, "layouts/main")
}
