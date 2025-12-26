package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

type R2Config struct {
	Bucket    string
	AccountID string
	Key       string
	Secret    string
}

type bucketBasic struct {
	Client *s3.Client
}

func NewR2() (*bucketBasic, error) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Failed to load env:", err)
	}

	r2 := R2Config{
		Bucket:    os.Getenv("BUCKET"),
		AccountID: os.Getenv("CL_ACCOUNT_ID"),
		Key:       os.Getenv("CL_ACCESS_KEY_ID"),
		Secret:    os.Getenv("CL_ACCESS_KEY"),
	}

	r2revolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", r2.AccountID),
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2revolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(r2.Key, r2.Secret, "")),
		config.WithRegion("auto"),
	)

	if err != nil {
		log.Fatal("Failed to initialize config:", err)
	}

	client := s3.NewFromConfig(cfg)

	basic := &bucketBasic{
		Client: client,
	}

	return basic, nil
}

func (b *bucketBasic) UploadFile(objectKey string, filename string, filetype string) error {
	file, err := os.Open(filename)
	bucketName := os.Getenv("BUCKET")

	if err != nil {
		log.Println("Couldn't open file: ", err)
		return err
	}

	defer file.Close()
	_, err = b.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(objectKey),
		Body:        file,
		ContentType: aws.String(filetype),
	})

	if err != nil {
		log.Println("Couldn't upload file to S3: ", err)
		return err
	}

	return nil
}

func (b *bucketBasic) Get(filename string) (string, error) {

	obj, err := b.Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET")),
		Key:    aws.String(filename),
	})
	defer obj.Body.Close()

	if err != nil {
		log.Println("Couldn't get object: ", err)

		return "", err
	}

	bodyBytes, err := io.ReadAll(obj.Body)

	return string(bodyBytes), nil
}

// func main() {

// 	err := godotenv.Load(".env")

// 	if err != nil {
// 		log.Fatal("Failed to load env:", err)
// 	}

// 	r2 := R2Config{
// 		Bucket:    os.Getenv("BUCKET"),
// 		AccountID: os.Getenv("CL_ACCOUNT_ID"),
// 		Key:       os.Getenv("CL_ACCESS_KEY_ID"),
// 		Secret:    os.Getenv("CL_ACCESS_KEY"),
// 	}

// 	r2revolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
// 		return aws.Endpoint{
// 			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", r2.AccountID),
// 		}, nil
// 	})

// 	cfg, err := config.LoadDefaultConfig(context.TODO(),
// 		config.WithEndpointResolverWithOptions(r2revolver),
// 		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(r2.Key, r2.Secret, "")),
// 		config.WithRegion("auto"),
// 	)

// 	if err != nil {
// 		log.Fatal("Failed to initialize config:", err)
// 	}

// 	// client := s3.NewFromConfig(cfg)

// 	// _ := bucketBasic{
// 	// 	Client: client,
// 	// }

// 	// listObjects, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
// 	// 	Bucket: &r2.Bucket,
// 	// })

// }
