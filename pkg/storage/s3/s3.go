package s3

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type ReplayStorage struct {
	Bucket    string
	Region    string
	AccessKey string
	SecretKey string
	Endpoint  string
}

func NewReplayStorage(bucket, region, accessKey, secretKey, endpoint string) ReplayStorage {
	return ReplayStorage{
		Bucket:    bucket,
		Region:    region,
		AccessKey: accessKey,
		SecretKey: secretKey,
		Endpoint:  endpoint,
	}
}

func (s ReplayStorage) Upload(gZipFilePath, target string) error {
	file, err := os.Open(gZipFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	cred := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(s.AccessKey, s.SecretKey, ""))
	s3Config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}

	client := s3.NewFromConfig(s3Config, func(o *s3.Options) {
		o.Credentials = cred
		o.BaseEndpoint = aws.String(s.Endpoint)
		o.Region = s.Region
		o.UsePathStyle = true
	})

	uploader := manager.NewUploader(client, func(u *manager.Uploader) {
		u.PartSize = 64 * 1024 * 1024 // 64MB per part
	})
	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(target),
		Body:   file,
	})

	return err
}

func (s ReplayStorage) TypeName() string {
	return "s3"
}
