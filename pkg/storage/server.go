package storage

import (
	"path/filepath"
	"strings"

	"github.com/jsthtlf/go-pam-sdk/model"
	"github.com/jsthtlf/go-pam-sdk/service"
)

type ServerStorage struct {
	pamService *service.PAMService
}

func (s ServerStorage) BulkSave(commands []*model.Command) error {
	return s.pamService.PushSessionCommand(commands)
}

func (s ServerStorage) Upload(gZipFilePath, target string) error {
	sessionID := strings.Split(filepath.Base(gZipFilePath), ".")[0]
	return s.pamService.Upload(sessionID, gZipFilePath)
}

func (s ServerStorage) TypeName() string {
	return StorageTypeServer
}
