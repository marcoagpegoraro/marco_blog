package mapper

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/marcoagpegoraro/marco_blog/dto"
	"github.com/marcoagpegoraro/marco_blog/models"
)

func MapPostRequestToPostModel(postRequest dto.PostRequest) models.Post {
	postModel := new(models.Post)

	postModel.Id = postRequest.PostId
	postModel.Title = postRequest.PostTitle
	postModel.Subtitle = postRequest.PostSubtitle
	postModel.Body = postRequest.PostBody

	var tags []models.Tag
	for _, tag := range strings.Split(postRequest.PostTags, ",") {
		tags = append(tags, models.Tag{Name: tag})
	}

	postModel.Tags = tags
	postModel.Language = postRequest.PostLanguage
	postModel.IsDraft = postRequest.PostIsDraft

	fmt.Println(postRequest.PostPublicatedAt)
	if postRequest.PostPublicatedAt == "" && postModel.IsDraft {
		postModel.PublicatedAt = sql.NullTime{}
	} else if postRequest.PostPublicatedAt == "" && !postModel.IsDraft {
		postModel.PublicatedAt = sql.NullTime{Time: time.Now(), Valid: true}
	} else {
		postPublicatedAt, err := time.Parse("2006-01-02 15:04:05 -0700 UTC", postRequest.PostPublicatedAt)
		if err != nil {
			fmt.Println(err)
		}
		postModel.PublicatedAt = sql.NullTime{Time: postPublicatedAt, Valid: true}
	}

	fmt.Println(postModel.PublicatedAt)
	return *postModel
}
