package aws_driver

import (
	"api-loyalty-point-agent/utils"
	"context"
	"log"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

const (
	bucket = "capstone14"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func UploadFileToBucket(filename string, src multipart.File) (string, error) {
	LoadEnv()

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Fatalf("error: %v\n", err)
		return "", err
	}

	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)

	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(utils.GenerateUniqueFileName(filename)),
		Body:   src,
	})

	log.Println("upload file to bucket")

	return result.Location, nil
}

func ReadAllFilesFromBucket() ([]string, error) {
	LoadEnv()

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Fatalf("error: %v\n", err)
		return []string{}, err
	}

	client := s3.NewFromConfig(cfg)
	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	})

	if err != nil {
		log.Fatalf("error: %v\n", err)
		return []string{}, err
	}

	log.Println("read data from bucket")

	var key []string
	for _, object := range output.Contents {
		key = append(key, aws.ToString(object.Key))
	}

	return key, nil
}

func DownloadFileFromBucket(filename string, dst string) error {
	LoadEnv()

	destination := dst + filename

	file, err := os.Create(destination)
	if err != nil {
		log.Fatalf("error: %v\n", err)
		return err
	}

	defer file.Close()

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Fatalf("error: %v\n", err)
		return err
	}

	client := s3.NewFromConfig(cfg)
	downloader := manager.NewDownloader(client)

	numBytes, err := downloader.Download(context.TODO(), file, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})

	if err != nil {
		log.Fatalf("error: %v\n", err)
		return err
	}

	log.Println("downloaded", file.Name(), numBytes, "bytes")

	return nil
}
