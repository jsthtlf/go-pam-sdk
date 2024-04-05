package http

import (
	"fmt"

	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

func (p *httpProvider) GetApplicationById(appId string) (app model.Application, err error) {
	reqUrl := fmt.Sprintf(UrlAppDetail, appId)
	_, err = p.authClient.Get(reqUrl, &app)
	return
}

func (p *httpProvider) GetRemoteApplicationById(remoteAppId string) (remoteApp model.RemoteAPP, err error) {
	Url := fmt.Sprintf(UrlRemoteAppDetail, remoteAppId)
	_, err = p.authClient.Get(Url, &remoteApp)
	return
}
