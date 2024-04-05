package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type ReplayStorage struct {
	Endpoint  string
	Bucket    string
	AccessKey string
	SecretKey string
}

func NewReplayStorage(endpoint, bucket, accessKey, secretKey string) ReplayStorage {
	return ReplayStorage{
		Endpoint:  endpoint,
		Bucket:    bucket,
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
}

func (o ReplayStorage) Upload(gZipFilePath, target string) error {
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

func (o ReplayStorage) TypeName() string {
	return "oss"
}
