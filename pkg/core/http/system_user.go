package http

import (
	"fmt"

	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

func (p *httpProvider) GetSystemUserById(systemUserId string) (sysUser model.SystemUser, err error) {
	url := fmt.Sprintf(UrlSystemUserDetail, systemUserId)
	_, err = p.get(url, &sysUser)
	return
}

func (p *httpProvider) GetUserApplicationAuthInfo(systemUserID, appID, userID, username string) (info model.SystemUserAuthInfo, err error) {
	Url := fmt.Sprintf(UrlSystemUserAppAuth, systemUserID, appID)
	params := make(map[string]string)
	if username != "" {
		params["username"] = username
	}
	if userID != "" {
		params["user_id"] = userID
	}
	_, err = p.get(Url, &info, params)
	return
}

func (p *httpProvider) GetUserApplicationSystemUsers(userId, appId string) (res []model.SystemUser, err error) {
	reqUrl := fmt.Sprintf(UrlUserPermsAppSystemUsers, userId, appId)
	_, err = p.get(reqUrl, &res)
	return
}

func (p *httpProvider) GetSystemUserAuthById(systemUserId, assetId, userId, username string) (info model.SystemUserAuthInfo, err error) {
	url := fmt.Sprintf(UrlSystemUserAuth, systemUserId)
	if assetId != "" {
		url = fmt.Sprintf(UrlSystemUserAssetAuth, systemUserId, assetId)
	}
	params := map[string]string{
		"username": username,
		"user_id":  userId,
	}
	_, err = p.get(url, &info, params)
	return
}
