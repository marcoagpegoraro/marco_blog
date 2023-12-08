package services

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/marcoagpegoraro/marco_blog/dto"
	"github.com/marcoagpegoraro/marco_blog/initializers"
	"github.com/marcoagpegoraro/marco_blog/mapper"
	"github.com/marcoagpegoraro/marco_blog/models"
	"github.com/marcoagpegoraro/marco_blog/repositories"
	"github.com/patrickmn/go-cache"
)

var PostService = PostServiceStruct{}

type PostServiceStruct struct {
}

func (service PostServiceStruct) GetIdParamFromUrl(c *fiber.Ctx) (int, error) {
	idParams, ok := c.AllParams()["id"]

	if !ok {
		return 0, fmt.Errorf("id not found in the url parameter")
	}

	lastOcurrenceOfDash := strings.LastIndex(idParams, "-")

	var idString string

	if lastOcurrenceOfDash == -1 {
		idString = idParams
	} else {
		idString = idParams[lastOcurrenceOfDash+1:]
	}

	idInt, err := strconv.Atoi(idString)

	if err != nil {
		return 0, err
	}

	return idInt, nil
}

func (service PostServiceStruct) HandlePostRequestPost(c *fiber.Ctx) (models.Post, error) {
	post := new(dto.PostRequest)

	if err := c.BodyParser(post); err != nil {
		fmt.Println("error = ", err)
		return models.Post{}, c.SendStatus(200)
	}

	imagesBase64 := AWSS3Service.GetBase64ImagesFromString(post.PostBody)
	if imagesBase64 != nil {
		imagesS3Url := AWSS3Service.UploadPostImagesToS3(imagesBase64)
		for index, imageBase64 := range imagesBase64 {
			post.PostBody = strings.Replace(post.PostBody, imageBase64, imagesS3Url[index], 1)
		}
	}

	postModel := mapper.MapPostRequestToPostModel(*post)
	return postModel, nil
}

func (service PostServiceStruct) GetTagsToBeDeleted(id int, newPost models.Post) []models.Tag {
	oldPost := repositories.PostRepository.GetPostById(id)

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

func (service PostServiceStruct) UpdatePost(post models.Post) {
	tagsToBeDeleted := service.GetTagsToBeDeleted(int(post.Id), post)

	repositories.TagRepository.DeleteTagsFromPost(&post, tagsToBeDeleted)
	repositories.PostRepository.Update(&post)
}

func (service PostServiceStruct) GetPostById(id int) models.Post {
	cacheKeyPost := fmt.Sprintf("postsServiceGetOne%d", id)

	var post models.Post
	if x, found := initializers.Cache.Get(cacheKeyPost); found {
		post = *x.(*models.Post)
	} else {
		post = repositories.PostRepository.GetPostById(id)
		initializers.Cache.Set(cacheKeyPost, &post, cache.DefaultExpiration)
	}

	return post
}
