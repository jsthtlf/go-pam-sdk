package service

import (
	"fmt"

	"github.com/jsthtlf/go-pam-sdk/model"
)

func (s *PAMService) GetAssetById(assetId string) (asset model.Asset, err error) {
	url := fmt.Sprintf(AssetDetailURL, assetId)
	_, err = s.authClient.Get(url, &asset)
	return
}

func (s *PAMService) GetAssetPlatform(assetId string) (platform model.Platform, err error) {
	url := fmt.Sprintf(AssetPlatFormURL, assetId)
	_, err = s.authClient.Get(url, &platform)
	return
}

func (s *PAMService) GetDomainGateways(domainId string) (domain model.Domain, err error) {
	Url := fmt.Sprintf(DomainDetailWithGateways, domainId)
	_, err = s.authClient.Get(Url, &domain)
	return
}

func (s *PAMService) GetSystemUserById(systemUserId string) (sysUser model.SystemUser, err error) {
	url := fmt.Sprintf(SystemUserDetailURL, systemUserId)
	_, err = s.authClient.Get(url, &sysUser)
	return
}

func (s *PAMService) GetSystemUserAuthById(systemUserId, assetId, userId,
	username string) (sysUser model.SystemUserAuthInfo, err error) {
	url := fmt.Sprintf(SystemUserAuthURL, systemUserId)
	if assetId != "" {
		url = fmt.Sprintf(SystemUserAssetAuthURL, systemUserId, assetId)
	}
	params := map[string]string{
		"username": username,
		"user_id":  userId,
	}
	_, err = s.authClient.Get(url, &sysUser, params)
	return
}
