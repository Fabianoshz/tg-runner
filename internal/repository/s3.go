package repository

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type S3 struct {
}

func NewPersistenceClient() Persistence {
	return &S3{}
}

func (a *S3) SavePlanfile(id string, planfile string, uuid uuid.UUID, path string) {
	client := s3.NewFromConfig(getConfig(), func(o *s3.Options) {
		if os.Getenv("DEV") == "true" {
			o.UsePathStyle = true
		}
	})

	body, file_err := os.ReadFile(path + "/" + planfile)
	if file_err != nil {
		fmt.Println("Got an error reading file: ", file_err)
	}

	input := &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String("plans/" + uuid.String() + "/" + id + "/planfile"),
		Body:   bytes.NewReader(body),
	}

	_, err := client.PutObject(context.TODO(), input)

	if err != nil {
		fmt.Println("Got an error: ", err)
	}
}

func (a *S3) GetPlanfiles(uuid uuid.UUID) {
	client := s3.NewFromConfig(getConfig(), func(o *s3.Options) {
		if os.Getenv("DEV") == "true" {
			o.UsePathStyle = true
		}
	})

	input := &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String("/plans/" + uuid.String()),
	}

	_, err := client.GetObject(context.TODO(), input)

	if err != nil {
		fmt.Println("Got an error: ", err)
	}
}

// TODO move this to internal/app/app.go
func getConfig() aws.Config {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(os.Getenv("AWS_REGION")),
	)

	if os.Getenv("DEV") == "true" {
		cfg.EndpointResolverWithOptions = aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL: os.Getenv("AWS_ENDPOINT"),
				}, nil
			},
		)
	}

	if err != nil {
		panic("configuration error, " + err.Error())
	}

	return cfg
}
