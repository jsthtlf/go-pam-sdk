package service

import (
	"fmt"

	"github.com/jsthtlf/go-pam-sdk/model"
)

func (s *PAMService) GetTokenAsset(token string) (tokenUser model.TokenUser, err error) {
	Url := fmt.Sprintf(TokenAssetURL, token)
	_, err = s.authClient.Get(Url, &tokenUser)
	return
}

func (s *PAMService) GetConnectTokenAuth(token string) (resp TokenAuthInfoResponse, err error) {
	data := map[string]string{
		"token": token,
	}
	_, err = s.authClient.Post(TokenAuthInfoURL, data, &resp)
	return
}

func (s *PAMService) RenewalToken(token string) (resp TokenRenewalResponse, err error) {
	data := map[string]string{
		"token": token,
	}
	_, err = s.authClient.Patch(TokenRenewalURL, data, &resp)
	return
}

type TokenRenewalResponse struct {
	Ok  bool   `json:"ok"`
	Msg string `json:"msg"`
}

type TokenAuthInfoResponse struct {
	Info model.ConnectTokenInfo
	Err  []string
}
