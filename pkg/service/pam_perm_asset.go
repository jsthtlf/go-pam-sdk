package service

import (
	"fmt"

	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

func (s *PAMService) SearchPermAsset(userId, key string) (res model.AssetList, err error) {
	Url := fmt.Sprintf(UserPermsAssetsURL, userId)
	payload := map[string]string{"search": key}
	_, err = s.authClient.Get(Url, &res, payload)
	return
}

func (s *PAMService) GetSystemUsersByUserIdAndAssetId(userId, assetId string) (sysUsers []model.SystemUser, err error) {
	Url := fmt.Sprintf(UserPermsAssetSystemUsersURL, userId, assetId)
	_, err = s.authClient.Get(Url, &sysUsers)
	return
}

func (s *PAMService) GetAllUserPermsAssets(userId string) ([]map[string]interface{}, error) {
	var params model.PaginationParam
	res, err := s.GetUserPermsAssets(userId, params)
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}

func (s *PAMService) GetUserPermsAssets(userID string, params model.PaginationParam) (resp model.PaginationResponse, err error) {
	Url := fmt.Sprintf(UserPermsAssetsURL, userID)
	return s.getPaginationResult(Url, params)
}

func (s *PAMService) RefreshUserAllPermsAssets(userId string) ([]map[string]interface{}, error) {
	var params model.PaginationParam
	params.Refresh = true
	res, err := s.GetUserPermsAssets(userId, params)
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}

func (s *PAMService) GetUserAssetByID(userId, assetId string) (assets []model.Asset, err error) {
	params := map[string]string{
		"id": assetId,
	}
	Url := fmt.Sprintf(UserPermsAssetsURL, userId)
	_, err = s.authClient.Get(Url, &assets, params)
	return
}

func (s *PAMService) GetUserPermAssetsByIP(userId, assetIP string) (assets []model.Asset, err error) {
	params := map[string]string{
		"ip": assetIP,
	}
	reqUrl := fmt.Sprintf(UserPermsAssetsURL, userId)
	_, err = s.authClient.Get(reqUrl, &assets, params)
	return
}