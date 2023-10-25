package services

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

var PostService = PostServiceStruct{}

type PostServiceStruct struct {
}

func (service PostServiceStruct) GetIdParamFromUrl(c *fiber.Ctx) (string, error) {
	params := c.AllParams()

	id, ok := params["id"]

	if !ok {
		return "", c.Render("pages/index/index", fiber.Map{
			"title": "Home",
		}, "layouts/main")
	}
	return id, nil
}

func (service PostServiceStruct) HandlePostRequestPost(c *fiber.Ctx) (models.Post, error) {
	post := new(dto.PostRequest)

	if err := c.BodyParser(post); err != nil {
		fmt.Println("error = ", err)
		return models.Post{}, c.SendStatus(200)
	}

	imagesBase64 := helpers.GetImagesFromString(post.PostBody)
	if imagesBase64 != nil {
		imagesS3Url := helpers.UploadPostImagesToS3(imagesBase64)
		for index, imageBase64 := range imagesBase64 {
			post.PostBody = strings.Replace(post.PostBody, imageBase64, imagesS3Url[index], 1)
		}
	}

	postModel := mapper.MapPostRequestToPostModel(*post)
	return postModel, nil
}

func (service PostServiceStruct) GetTagsToBeDeleted(id uint, newPost models.Post) []models.Tag {
	var oldPost models.Post
	initializers.DB.Where("id = ?", id).Preload("Tags").First(&oldPost)

	var tagsToBeDeleted []models.Tag

	for _, oldTag := range oldPost.Tags {
		hasTagInNewAndOldModel := false
		for _, newTag := range newPost.Tags {
			if oldTag.Name == newTag.Name {
				hasTagInNewAndOldModel = true
				break
			}
		}

		if !hasTagInNewAndOldModel {
			tagsToBeDeleted = append(tagsToBeDeleted, oldTag)
		}
	}

	return tagsToBeDeleted
}
