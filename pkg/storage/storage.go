package storage

import (
	"strings"

	"github.com/jsthtlf/go-pam-sdk/pkg/model"
	"github.com/jsthtlf/go-pam-sdk/pkg/storage/azure"
	"github.com/jsthtlf/go-pam-sdk/pkg/storage/elasticsearch"
	"github.com/jsthtlf/go-pam-sdk/pkg/storage/obs"
	"github.com/jsthtlf/go-pam-sdk/pkg/storage/oss"
	"github.com/jsthtlf/go-pam-sdk/pkg/storage/s3"
)

const (
	TypeAzure = "azure"

	TypeOSS = "oss"

	TypeS3    = "s3"
	TypeSwift = "swift"
	TypeCOS   = "cos"

	TypeOBS = "obs"

	TypeNull = "null"

	TypeServer = "server"

	TypeES            = "es"
	TypeElasticSearch = "elasticsearch"
)

type Typer interface {
	TypeName() string
}

type CommandStorage interface {
	Typer

	BulkSave(commands []*model.Command) error
}

type ReplayStorage interface {
	Typer

	Upload(gZipFile, target string) error
}

func NewCommandStorage(p serverProvider, conf model.CommandConfig) CommandStorage {
	switch conf.TypeName {
	case TypeES, TypeElasticSearch:
		hosts := conf.Hosts
		index := conf.Index
		docType := conf.DocType
		skipVerify := conf.Other.IgnoreVerifyCerts
		if index == "" {
			index = "pam"
		}
		if docType == "" {
			docType = "_doc"
		}
		return elasticsearch.NewCommandStorage(hosts, index, docType, skipVerify)

	case TypeNull:
		return NewNullStorage()

	default:
		return NewStorage(p)
	}
}

func NewReplayStorage(p serverProvider, conf model.ReplayConfig) ReplayStorage {
	switch conf.TypeName {
	case TypeAzure:
		accountName := conf.AccountName
		accountKey := conf.AccountKey
		containerName := conf.ContainerName
		endpointSuffix := conf.EndpointSuffix

		if endpointSuffix == "" {
			endpointSuffix = "core.chinacloudapi.cn"
		}

		return azure.NewReplayStorage(accountName, accountKey, containerName, endpointSuffix)

	case TypeOSS:
		endpoint := conf.Endpoint
		bucket := conf.Bucket
		accessKey := conf.AccessKey
		secretKey := conf.SecretKey

		return oss.NewReplayStorage(endpoint, bucket, accessKey, secretKey)

	case TypeS3, TypeSwift, TypeCOS:
		bucket := conf.Bucket
		endpoint := conf.Endpoint
		region := conf.Region
		accessKey := conf.AccessKey
		secretKey := conf.SecretKey

		if region == "" && endpoint != "" {
			endpointArray := strings.Split(endpoint, ".")
			if len(endpointArray) >= 2 {
				region = endpointArray[1]
			}
		}
		if bucket == "" {
			bucket = "pamservice"
		}

		return s3.NewReplayStorage(bucket, region, accessKey, secretKey, endpoint)

	case TypeOBS:
		endpoint := conf.Endpoint
		bucket := conf.Bucket
		accessKey := conf.AccessKey
		secretKey := conf.SecretKey

		return obs.NewReplayStorage(endpoint, bucket, accessKey, secretKey)

	case TypeNull:
		return NewNullStorage()

	default:
		return NewStorage(p)
	}
}
