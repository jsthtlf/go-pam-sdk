package storage

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3ReplayStorage struct {
	Bucket    string
	Region    string
	AccessKey string
	SecretKey string
	Endpoint  string
}

func (s S3ReplayStorage) Upload(gZipFilePath, target string) error {
	file, err := os.Open(gZipFilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(s.AccessKey, s.SecretKey, ""),
		Endpoint:         aws.String(s.Endpoint),
		Region:           aws.String(s.Region),
		S3ForcePathStyle: aws.Bool(true),
	}

	sess, err := session.NewSession(s3Config)
	if err != nil {
		return err
	}

	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		u.PartSize = 64 * 1024 * 1024 // 64MB per part
	})
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(target),
		Body:   file,
	})

	return err
}

func (s S3ReplayStorage) TypeName() string {
	return StorageTypeS3
}