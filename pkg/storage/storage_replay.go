package storage

import (
	"github.com/jsthtlf/go-pam-sdk/model"
	"github.com/jsthtlf/go-pam-sdk/service"
	"strings"
)

type ReplayStorage interface {
	Upload(gZipFile, target string) error
	StorageType
}

func NewReplayStorage(pamService *service.PAMService, conf *model.TerminalConfig) ReplayStorage {
	cfg := conf.ReplayStorage
	switch cfg.TypeName {
	case StorageTypeAzure:
		var (
			accountName    string
			accountKey     string
			containerName  string
			endpointSuffix string
		)
		endpointSuffix = cfg.EndpointSuffix
		accountName = cfg.AccountName
		accountKey = cfg.AccountKey
		containerName = cfg.ContainerName
		if endpointSuffix == "" {
			endpointSuffix = "core.chinacloudapi.cn"
		}
		return AzureReplayStorage{
			AccountName:    accountName,
			AccountKey:     accountKey,
			ContainerName:  containerName,
			EndpointSuffix: endpointSuffix,
		}
	case StorageTypeOSS:
		var (
			endpoint  string
			bucket    string
			accessKey string
			secretKey string
		)

		endpoint = cfg.Endpoint
		bucket = cfg.Bucket
		accessKey = cfg.AccessKey
		secretKey = cfg.SecretKey

		return OSSReplayStorage{
			Endpoint:  endpoint,
			Bucket:    bucket,
			AccessKey: accessKey,
			SecretKey: secretKey,
		}
	case StorageTypeS3, StorageTypeSwift, StorageTypeCOS:
		var (
			region    string
			endpoint  string
			bucket    string
			accessKey string
			secretKey string
		)

		bucket = cfg.Bucket
		endpoint = cfg.Endpoint
		region = cfg.Region
		accessKey = cfg.AccessKey
		secretKey = cfg.SecretKey

		if region == "" && endpoint != "" {
			endpointArray := strings.Split(endpoint, ".")
			if len(endpointArray) >= 2 {
				region = endpointArray[1]
			}
		}
		if bucket == "" {
			bucket = "pamservice"
		}
		return S3ReplayStorage{
			Bucket:    bucket,
			Region:    region,
			AccessKey: accessKey,
			SecretKey: secretKey,
			Endpoint:  endpoint,
		}
	case StorageTypeOBS:
		var (
			endpoint  string
			bucket    string
			accessKey string
			secretKey string
		)

		endpoint = cfg.Endpoint
		bucket = cfg.Bucket
		accessKey = cfg.AccessKey
		secretKey = cfg.SecretKey

		return OBSReplayStorage{
			Endpoint:  endpoint,
			Bucket:    bucket,
			AccessKey: accessKey,
			SecretKey: secretKey,
		}
	case StorageTypeNull:
		return NewNullStorage()
	default:
		return ServerStorage{pamService: pamService}
	}
}
