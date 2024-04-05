package http

import (
	"fmt"

	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

func (p *httpProvider) GetTokenAsset(token string) (tokenUser model.TokenUser, err error) {
	Url := fmt.Sprintf(UrlTokenAsset, token)
	_, err = p.authClient.Get(Url, &tokenUser)
	return
}

func (p *httpProvider) GetConnectTokenAuth(token string) (resp model.ConnectTokenInfo, err error) {
	data := map[string]string{
		"token": token,
	}
	_, err = p.authClient.Post(UrlTokenAuthInfo, data, &resp)
	return
}

func (p *httpProvider) RenewalToken(token string) (resp model.TokenRenewalResponse, err error) {
	data := map[string]string{
		"token": token,
	}
	_, err = p.authClient.Patch(UrlTokenRenewal, data, &resp)
	return
}
