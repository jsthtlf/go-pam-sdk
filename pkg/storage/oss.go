package storage

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type OSSReplayStorage struct {
	Endpoint  string
	Bucket    string
	AccessKey string
	SecretKey string
}

func NewOSSReplayStorage(endpoint, bucket, accessKey, secretKey string) OSSReplayStorage {
	return OSSReplayStorage{
		Endpoint:  endpoint,
		Bucket:    bucket,
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
}

func (o OSSReplayStorage) Upload(gZipFilePath, target string) error {
	client, err := oss.New(o.Endpoint, o.AccessKey, o.SecretKey)
	if err != nil {
		return err
	}
	bucket, err := client.Bucket(o.Bucket)
	if err != nil {
		return err
	}
	return bucket.PutObjectFromFile(target, gZipFilePath)
}

func (o OSSReplayStorage) TypeName() string {
	return TypeOSS
}
