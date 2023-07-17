package storage

import (
	"path/filepath"
	"strings"

	"github.com/jsthtlf/go-pam-sdk/model"
	"github.com/jsthtlf/go-pam-sdk/service"
)

type ServerStorage struct {
	StorageType string
	PAMService  *service.PAMService
}

func (s ServerStorage) BulkSave(commands []*model.Command) (err error) {
	return s.PAMService.PushSessionCommand(commands)
}

func (s ServerStorage) Upload(gZipFilePath, target string) (err error) {
	sessionID := strings.Split(filepath.Base(gZipFilePath), ".")[0]
	return s.PAMService.Upload(sessionID, gZipFilePath)
}

func (s ServerStorage) TypeName() string {
	return s.StorageType
}
