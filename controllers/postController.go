package controllers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/dto"
	"github.com/marcoagpegoraro/marco_blog/helpers"
	"github.com/marcoagpegoraro/marco_blog/initializers"
	"github.com/marcoagpegoraro/marco_blog/mapper"
	"github.com/marcoagpegoraro/marco_blog/models"
)

func GetPostIndex(c *fiber.Ctx) error {
	return c.Render("posts/index", fiber.Map{
		"Title": "Create new post",
	}, "layouts/main")
}

func GetOnePostIndex(c *fiber.Ctx) error {
	params := c.AllParams()

	id, ok := params["id"]

	if !ok {
		return c.Render("index/index", fiber.Map{
			"Title": "Home",
		}, "layouts/main")
	}

	var post models.Post
	initializers.DB.Where("id = ?", id).First(&post)

	return c.Render("posts/one", fiber.Map{
		"Title": "Create new post",
		"Post":  post,
	}, "layouts/main")
}

func PostPostIndex(c *fiber.Ctx) error {
	post := new(dto.PostRequest)

	if err := c.BodyParser(post); err != nil {
		fmt.Println("error = ", err)
		return c.SendStatus(200)
	}

	imagesBase64 := helpers.GetImagesFromString(post.PostBody)
	imagesS3Url := helpers.UploadPostImagesToS3(imagesBase64)

	for index, imageBase64 := range imagesBase64 {
		post.PostBody = strings.Replace(post.PostBody, imageBase64, imagesS3Url[index], 1)
	}

	postModel := mapper.MapPostRequestToPostModel(*post)
	initializers.DB.Create(&postModel)

	return c.Render("posts/index", fiber.Map{
		"Title": "Create new post",
	}, "layouts/main")
}
