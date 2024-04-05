package obs

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
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

func (o ReplayStorage) TypeName() string {
	return "obs"
}
