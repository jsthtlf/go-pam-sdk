package service

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/jsthtlf/go-pam-sdk/pkg/model"
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
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expired_at"`
}

type TokenAuthInfoResponse struct {
	Info model.ConnectTokenInfo
	Err  []string
}

func (t *TokenAuthInfoResponse) UnmarshalJSON(p []byte) error {
	if index := bytes.IndexByte(p, '['); index == 0 {
		return json.Unmarshal(p, &t.Err)
	}
	return json.Unmarshal(p, &t.Info)
}
