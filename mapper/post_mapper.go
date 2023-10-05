package mapper

import (
	"strings"

	"github.com/marcoagpegoraro/marco_blog/dto"
	"github.com/marcoagpegoraro/marco_blog/models"
)

func MapPostRequestToPostModel(postRequest dto.PostRequest) models.Post {
	postModel := new(models.Post)

	postModel.Title = postRequest.PostTitle
	postModel.Subtitle = postRequest.PostSubtitle
	postModel.Body = postRequest.PostBody

	var tags []models.Tag
	for _, tag := range strings.Split(postRequest.PostTags, ",") {
		tags = append(tags, models.Tag{Name: tag})
	}

	postModel.Tags = tags

	return *postModel
}
