package helpers_test

import (
	"log"
	"os"
	"testing"

	"github.com/marcoagpegoraro/marco_blog/helpers"
)

func TestGetImageFromString(t *testing.T) {

	body, err := os.ReadFile("./test_files/post_body.html")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	postBody := string(body)

	images := helpers.GetImagesFromString(postBody)

	if images == nil || len(images) != 3 {
		t.Errorf("Lenght of images correct")
	}

	for _, image := range images {
		if image[0:4] != "data" {
			t.Errorf("base64 image not valid")
		}
	}
}
