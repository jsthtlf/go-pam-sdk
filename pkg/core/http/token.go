package http

import (
	"fmt"

	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

func (p *httpProvider) GetTokenAsset(token string) (tokenUser model.TokenUser, err error) {
	Url := fmt.Sprintf(UrlTokenAsset, token)
	_, err = p.get(Url, &tokenUser)
	return
}

func (p *httpProvider) GetConnectTokenAuth(token string) (resp model.ConnectTokenInfo, err error) {
	data := map[string]string{
		"token": token,
	}
	_, err = p.post(UrlTokenAuthInfo, data, &resp)
	return
}

func (p *httpProvider) RenewalToken(token string) (resp model.TokenRenewalResponse, err error) {
	data := map[string]string{
		"token": token,
	}
	_, err = p.patch(UrlTokenRenewal, data, &resp)
	return
}
