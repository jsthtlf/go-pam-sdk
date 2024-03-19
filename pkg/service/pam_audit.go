package service

import (
	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

func (s *PAMService) CreateFileOperationLog(data model.FTPLog) (err error) {
	_, err = s.authClient.Post(FTPLogListURL, data, nil)
	return
}

func (s *PAMService) PushSessionCommand(commands []*model.Command) (err error) {
	_, err = s.authClient.Post(SessionCommandURL, commands, nil)
	return
}

func (s *PAMService) NotifyCommand(commands []*model.Command) (err error) {
	_, err = s.authClient.Post(NotificationCommandURL, commands, nil)
	return
}
