package http

import (
	"fmt"

	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

func (p *httpProvider) GetAssetById(assetId string) (asset model.Asset, err error) {
	url := fmt.Sprintf(UrlAssetDetail, assetId)
	_, err = p.authClient.Get(url, &asset)
	return
}

func (p *httpProvider) GetAssetPlatform(assetId string) (platform model.Platform, err error) {
	url := fmt.Sprintf(UrlAssetPlatFormDetail, assetId)
	_, err = p.authClient.Get(url, &platform)
	return
}

func (p *httpProvider) GetDomainGateways(domainId string) (domain model.Domain, err error) {
	Url := fmt.Sprintf(UrlAssetDomainDetail, domainId)
	_, err = p.authClient.Get(Url, &domain)
	return
}
