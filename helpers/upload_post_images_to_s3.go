package helpers

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/google/uuid"
	"github.com/marcoagpegoraro/marco_blog/initializers"
)

var bucketName = "marco-blog-post-images"
var s3UrlPrefix = "https://" + bucketName + ".s3.sa-east-1.amazonaws.com/"

func UploadPostImagesToS3(images []string) []string {
	var imagesUrlS3 []string

	for _, image := range images {

		b64data := image[strings.IndexByte(image, ',')+1:]

		decode, err := base64.StdEncoding.DecodeString(b64data)

		if err != nil {
			return nil
		}

		imageName := uuid.New().String() + ".jpg"

		imagesUrlS3 = append(imagesUrlS3, s3UrlPrefix+imageName)

		obj, err := initializers.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(imageName),
			Body:   bytes.NewReader(decode),
			ACL:    types.ObjectCannedACLPublicRead,
		})

		if err != nil {
			fmt.Println("Error sending image to s3")
			fmt.Println(err)
			fmt.Println(obj)
		}

	}
	return imagesUrlS3
}

func GetImagesFromString(str string) []string {
	pattern := regexp.MustCompile("<img src=\"(.*?)\">")
	matches := pattern.FindAllStringSubmatch(str, -1)

	if matches == nil {
		return nil
	}

	var base64Images []string

	for _, v := range matches {
		if v[1][0:4] == "data" {
			base64Images = append(base64Images, v[1])
		}
	}

	matches = nil

	return base64Images
}
