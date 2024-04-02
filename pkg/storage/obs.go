package storage

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
)

type OBSReplayStorage struct {
	Endpoint  string
	Bucket    string
	AccessKey string
	SecretKey string
}

func NewOBSReplayStorage(endpoint, bucket, accessKey, secretKey string) OBSReplayStorage {
	return OBSReplayStorage{
		Endpoint:  endpoint,
		Bucket:    bucket,
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
}

func (o OBSReplayStorage) Upload(gZipFilePath, target string) error {
	client, err := obs.New(o.AccessKey, o.SecretKey, o.Endpoint)
	if err != nil {
		return err
	}
	input := &obs.PutFileInput{}
	input.Bucket = o.Bucket
	input.Key = target
	input.SourceFile = gZipFilePath
	_, err = client.PutFile(input)

	return err
}

func (o OBSReplayStorage) TypeName() string {
	return TypeOBS
}
