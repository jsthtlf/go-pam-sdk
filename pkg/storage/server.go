package storage

import (
	"path/filepath"
	"strings"

	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

type serverProvider interface {
	CreateSessionCommand(commands []*model.Command) (err error)
	UploadReplay(sid, gZipFile string) error
}

type Storage struct {
	p serverProvider
}

func NewStorage(p serverProvider) Storage {
	return Storage{p: p}
}

func (s Storage) BulkSave(commands []*model.Command) error {
	return s.p.CreateSessionCommand(commands)
}

func (s Storage) Upload(gZipFilePath, target string) error {
	sessionID := strings.Split(filepath.Base(gZipFilePath), ".")[0]
	return s.p.UploadReplay(sessionID, gZipFilePath)
}

func (s Storage) TypeName() string {
	return "server"
}
