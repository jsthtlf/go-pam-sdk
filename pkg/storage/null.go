package storage

import (
	"github.com/jsthtlf/go-pam-sdk/pkg/logger"
	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

func NewNullStorage() NullStorage {
	return NullStorage{}
}

type NullStorage struct {
}

func (f NullStorage) BulkSave(commands []*model.Command) error {
	logger.Infof("Null Storage discard %d commands.", len(commands))
	return nil
}

func (f NullStorage) Upload(gZipFile, target string) error {
	logger.Infof("Null Storage discard %s.", gZipFile)
	return nil
}

func (f NullStorage) TypeName() string {
	return StorageTypeNull
}
