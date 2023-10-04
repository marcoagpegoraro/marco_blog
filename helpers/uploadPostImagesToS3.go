package helpers

import (
	"bytes"
	"context"
	"encoding/base64"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/marcoagpegoraro/marco_blog/initializers"
)

func UploadPostImagesToS3(postBody string) error {
	images := GetImagesFromString(postBody)

	for _, image := range images {

		b64data := image[strings.IndexByte(image, ',')+1:]

		decode, err := base64.StdEncoding.DecodeString(b64data)

		if err != nil {
			return err
		}

		initializers.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String("marco-blog-post-images"),
			Key:    aws.String(uuid.New().String() + ".jpg"),
			Body:   bytes.NewReader(decode),
		})

	}
	return nil
}

func GetImagesFromString(str string) []string {
	pattern := regexp.MustCompile("<img src=\"(.*?)\">")
	matches := pattern.FindAllStringSubmatch(str, -1)

	if matches == nil {
		return nil
	}

	var base64Images []string

	for _, v := range matches {
		base64Images = append(base64Images, v[1])
	}

	matches = nil

	return base64Images
}
