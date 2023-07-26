package storage

import (
	"strings"

	"github.com/jsthtlf/go-pam-sdk/model"
	"github.com/jsthtlf/go-pam-sdk/service"
)

type StorageType interface {
	TypeName() string
}

type ReplayStorage interface {
	Upload(gZipFile, target string) error
	StorageType
}

type CommandStorage interface {
	BulkSave(commands []*model.Command) error
	StorageType
}

func NewReplayStorage(pamService *service.PAMService, conf *model.TerminalConfig) ReplayStorage {
	cfg := conf.ReplayStorage
	switch cfg.TypeName {
	case "azure":
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
	case "oss":
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
	case "s3", "swift", "cos":
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
	case "obs":
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
	case "null":
		return NewNullStorage()
	default:
		return ServerStorage{StorageType: "server", PAMService: pamService}
	}
}

func NewCommandStorage(pamService *service.PAMService, conf *model.TerminalConfig) CommandStorage {
	cf := conf.CommandStorage
	tp, ok := cf["TYPE"]
	if !ok {
		tp = "server"
	}
	switch tp {
	case "es", "elasticsearch":
		var hosts = make([]string, len(cf["HOSTS"].([]interface{})))
		for i, item := range cf["HOSTS"].([]interface{}) {
			hosts[i] = item.(string)
		}
		var skipVerify bool
		index := cf["INDEX"].(string)
		docType := cf["DOC_TYPE"].(string)
		if otherMap, ok := cf["OTHER"].(map[string]interface{}); ok {
			if insecureSkipVerify, ok := otherMap["IGNORE_VERIFY_CERTS"]; ok {
				skipVerify = insecureSkipVerify.(bool)
			}
		}
		if index == "" {
			index = "pam"
		}
		if docType == "" {
			docType = "_doc"
		}
		return ESCommandStorage{
			Hosts:              hosts,
			Index:              index,
			DocType:            docType,
			InsecureSkipVerify: skipVerify,
		}
	case "null":
		return NewNullStorage()
	default:
		return ServerStorage{StorageType: "server", PAMService: pamService}
	}
}
