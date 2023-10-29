package mapper_test

import (
	"testing"

	"github.com/marcoagpegoraro/marco_blog/dto"
	"github.com/marcoagpegoraro/marco_blog/mapper"
)

func TestMapPostRequestToPostModel(t *testing.T) {
	postRequest := dto.PostRequest{
		PostPublicatedAt: "2023-10-27 21:44:06.531 +0000 UTC",
	}
	postModel := mapper.MapPostRequestToPostModel(postRequest)

	if postModel.PublicatedAt.Time.String() != "2023-10-27 21:44:06.531 +0000 +0000" {
		t.Errorf("Parse of publicated date is wrong")
	}
}
