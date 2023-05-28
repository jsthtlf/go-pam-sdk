package service

import (
	"fmt"
	"github.com/jsthtlf/go-pam-sdk/model"
)

func (s *PAMService) GetApplicationById(appId string) (app model.Application, err error) {
	reqUrl := fmt.Sprintf(ApplicationDetailURL, appId)
	_, err = s.authClient.Get(reqUrl, &app)
	return
}

func (s *PAMService) GetUserApplicationAuthInfo(systemUserID, appID, userID, username string) (info model.SystemUserAuthInfo, err error) {
	Url := fmt.Sprintf(SystemUserAppAuthURL, systemUserID, appID)
	params := make(map[string]string)
	if username != "" {
		params["username"] = username
	}
	if userID != "" {
		params["user_id"] = userID
	}
	_, err = s.authClient.Get(Url, &info, params)
	return
}

func (s *PAMService) GetUserApplicationSystemUsers(userId, appId string) (res []model.SystemUser, err error) {
	reqUrl := fmt.Sprintf(UserPermsApplicationSystemUsersURL, userId, appId)
	_, err = s.authClient.Get(reqUrl, &res)
	return
}
