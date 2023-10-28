package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/enum"
	"github.com/marcoagpegoraro/marco_blog/initializers"
	"github.com/marcoagpegoraro/marco_blog/models"
	"github.com/marcoagpegoraro/marco_blog/services"
	"github.com/patrickmn/go-cache"
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

	cacheKeyPost := fmt.Sprintf("postsControllerGetOne%s", id)

	var post models.Post
	if x, found := initializers.Cache.Get(cacheKeyPost); found {
		post = *x.(*models.Post)
	} else {
		initializers.DB.Where("id = ?", id).Preload("Tags").First(&post)
		initializers.Cache.Set(cacheKeyPost, &post, cache.DefaultExpiration)
	}

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
	postModel, err := services.PostService.HandlePostRequestPost(c)
	if err != nil {
		return err
	}

	initializers.DB.Create(&postModel)

	initializers.Cache.Flush()

	return c.Redirect(fmt.Sprintf("/posts/%d", postModel.Id))
}

func (controller PostControllerStruct) GetEditPost(c *fiber.Ctx) error {
	id, err := services.PostService.GetIdParamFromUrl(c)
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
	postModel, err := services.PostService.HandlePostRequestPost(c)
	if err != nil {
		return err
	}

	tagsToBeDeleted := services.PostService.GetTagsToBeDeleted(postModel.Id, postModel)

	initializers.DB.Model(&postModel).Association("Tags").Delete(tagsToBeDeleted)
	initializers.DB.Omit("created_at").Save(&postModel)

	initializers.Cache.Flush()

	return c.Redirect(fmt.Sprintf("/posts/%d", postModel.Id))
}
