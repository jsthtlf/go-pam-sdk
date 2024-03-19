package service

import (
	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

func (s *PAMService) CheckUserCookie(cookies map[string]string) (user *model.User, err error) {
	client := s.authClient.Clone()
	for k, v := range cookies {
		client.SetCookie(k, v)
	}
	_, err = client.Get(UserProfileURL, &user)
	return
}
