package initializers

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var S3Client *s3.Client

func ConnectToS3() {

	awsAccessKeyId := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	if awsAccessKeyId == "" || awsSecretAccessKey == "" {
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Printf("error: %v", err)
			return
		}
		S3Client = s3.NewFromConfig(cfg)
	} else {
		S3Client = s3.New(s3.Options{
			Region: "sa-east-1",
			Credentials: aws.NewCredentialsCache(
				credentials.NewStaticCredentialsProvider(awsAccessKeyId, awsSecretAccessKey, "")),
		})
	}

}
