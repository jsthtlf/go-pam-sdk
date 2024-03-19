package storage

import (
	"github.com/jsthtlf/go-pam-sdk/model"
	"github.com/jsthtlf/go-pam-sdk/service"
)

type CommandStorage interface {
	BulkSave(commands []*model.Command) error
	StorageType
}

func NewCommandStorage(pamService *service.PAMService, conf *model.TerminalConfig) CommandStorage {
	cf := conf.CommandStorage
	tp, ok := cf["TYPE"]
	if !ok {
		tp = StorageTypeServer
	}
	switch tp {
	case StorageTypeES, StorageTypeElasticSearch:
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
	case StorageTypeNull:
		return NewNullStorage()
	default:
		return ServerStorage{pamService: pamService}
	}
}
