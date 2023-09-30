package mapper

import (
	"github.com/marcoagpegoraro/marco_blog/dto"
	"github.com/marcoagpegoraro/marco_blog/models"
)

func MapPostRequestToPostModel(postRequest dto.PostRequest) models.Post {
	postModel := new(models.Post)

	postModel.Title = postRequest.PostTitle
	postModel.Subtitle = postRequest.PostSubtitle
	postModel.Body = postRequest.PostBody

	return *postModel
}
