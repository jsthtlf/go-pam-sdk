package service

import (
	"fmt"

	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

func (s *PAMService) GetRemoteApp(remoteAppId string) (remoteApp model.RemoteAPP, err error) {
	Url := fmt.Sprintf(RemoteAPPURL, remoteAppId)
	_, err = s.authClient.Get(Url, &remoteApp)
	return
}
