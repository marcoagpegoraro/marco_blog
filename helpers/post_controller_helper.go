package helpers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/dto"
	"github.com/marcoagpegoraro/marco_blog/mapper"
	"github.com/marcoagpegoraro/marco_blog/models"
)

func GetIdParamFromUrl(c *fiber.Ctx) (string, error) {
	params := c.AllParams()

	id, ok := params["id"]

	if !ok {
		return "", c.Render("index/index", fiber.Map{
			"title": "Home",
		}, "layouts/main")
	}
	return id, nil
}

func HandlePostRequestPost(c *fiber.Ctx) (models.Post, error) {
	post := new(dto.PostRequest)

	if err := c.BodyParser(post); err != nil {
		fmt.Println("error = ", err)
		return models.Post{}, c.SendStatus(200)
	}

	imagesBase64 := GetImagesFromString(post.PostBody)
	if imagesBase64 != nil {
		imagesS3Url := UploadPostImagesToS3(imagesBase64)
		for index, imageBase64 := range imagesBase64 {
			post.PostBody = strings.Replace(post.PostBody, imageBase64, imagesS3Url[index], 1)
		}
	}

	postModel := mapper.MapPostRequestToPostModel(*post)
	return postModel, nil
}
