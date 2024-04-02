package storage

import (
	"strings"

	"github.com/jsthtlf/go-pam-sdk/pkg/model"
	"github.com/jsthtlf/go-pam-sdk/pkg/service"
)

const (
	TypeAzure         = "azure"
	TypeOSS           = "oss"
	TypeS3            = "s3"
	TypeSwift         = "swift"
	TypeCOS           = "cos"
	TypeOBS           = "obs"
	TypeNull          = "null"
	TypeServer        = "server"
	TypeES            = "es"
	TypeElasticSearch = "elasticsearch"
)

type Typer interface {
	TypeName() string
}

type CommandStorage interface {
	BulkSave(commands []*model.Command) error
	Typer
}

type ReplayStorage interface {
	Upload(gZipFile, target string) error
	Typer
}

func NewCommandStorage(pamService *service.PAMService, conf *model.TerminalConfig) CommandStorage {
	cfg := conf.CommandStorage
	switch cfg.TypeName {
	case TypeES, TypeElasticSearch:
		hosts := cfg.Hosts
		index := cfg.Index
		docType := cfg.DocType
		skipVerify := cfg.Other.IgnoreVerifyCerts
		if index == "" {
			index = "pam"
		}
		if docType == "" {
			docType = "_doc"
		}
		return NewESCommandStorage(hosts, index, docType, skipVerify)

	case TypeNull:
		return NewNullStorage()

	default:
		return NewServerStorage(pamService)
	}
}

func NewReplayStorage(pamService *service.PAMService, conf *model.TerminalConfig) ReplayStorage {
	cfg := conf.ReplayStorage
	switch cfg.TypeName {
	case TypeAzure:
		accountName := cfg.AccountName
		accountKey := cfg.AccountKey
		containerName := cfg.ContainerName
		endpointSuffix := cfg.EndpointSuffix

		if endpointSuffix == "" {
			endpointSuffix = "core.chinacloudapi.cn"
		}

		return NewAzureReplayStorage(accountName, accountKey, containerName, endpointSuffix)

	case TypeOSS:
		endpoint := cfg.Endpoint
		bucket := cfg.Bucket
		accessKey := cfg.AccessKey
		secretKey := cfg.SecretKey

		return NewOSSReplayStorage(endpoint, bucket, accessKey, secretKey)

	case TypeS3, TypeSwift, TypeCOS:
		bucket := cfg.Bucket
		endpoint := cfg.Endpoint
		region := cfg.Region
		accessKey := cfg.AccessKey
		secretKey := cfg.SecretKey

		if region == "" && endpoint != "" {
			endpointArray := strings.Split(endpoint, ".")
			if len(endpointArray) >= 2 {
				region = endpointArray[1]
			}
		}
		if bucket == "" {
			bucket = "pamservice"
		}

		return NewS3ReplayStorage(bucket, region, accessKey, secretKey, endpoint)

	case TypeOBS:
		endpoint := cfg.Endpoint
		bucket := cfg.Bucket
		accessKey := cfg.AccessKey
		secretKey := cfg.SecretKey

		return NewOBSReplayStorage(endpoint, bucket, accessKey, secretKey)

	case TypeNull:
		return NewNullStorage()

	default:
		return NewServerStorage(pamService)
	}
}
