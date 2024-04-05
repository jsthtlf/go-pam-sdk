package http

import (
	"fmt"
	"net/http"

	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

func (p *httpProvider) GetUserById(uid string) (user *model.User, err error) {
	url := fmt.Sprintf(UrlUserDetail, uid)
	_, err = p.authClient.Get(url, &user)
	return
}

func (p *httpProvider) GetProfile() (user *model.User, err error) {
	var res *http.Response
	res, err = p.authClient.Get(UrlUserProfile, &user)
	if res == nil && err != nil {
		return nil, fmt.Errorf("%w:%v", ErrConnect, err)
	}
	if res != nil && res.StatusCode == http.StatusUnauthorized {
		return user, ErrAccessKeyUnauthorized
	}
	return user, err
}

func (p *httpProvider) CheckUserCookie(cookies map[string]string) (user *model.User, err error) {
	client := p.authClient.Clone()
	for k, v := range cookies {
		client.SetCookie(k, v)
	}
	_, err = client.Get(UrlUserProfile, &user)
	return
}

func (p *httpProvider) GetShareUserInfo(query string) (res []*model.MiniUser, err error) {
	params := make(map[string]string)
	params["action"] = "suggestion"
	params["search"] = query
	_, err = p.authClient.Get(UrlUserList, &res, params)
	return
}
