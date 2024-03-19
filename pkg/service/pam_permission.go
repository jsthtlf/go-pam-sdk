package service

import (
	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

func (s *PAMService) GetPermission(userId, assetId, systemUserId string) (perms model.Permission, err error) {
	params := map[string]string{
		"user_id":        userId,
		"asset_id":       assetId,
		"system_user_id": systemUserId,
	}
	_, err = s.authClient.Get(PermissionURL, &perms, params)
	return
}

func (s *PAMService) ValidateRemoteAppPermission(userId, remoteAppId, systemUserId string) (info model.ExpireInfo, err error) {
	return s.ValidateApplicationPermission(userId, remoteAppId, systemUserId)
}

func (s *PAMService) ValidateApplicationPermission(userId, appId, systemUserId string) (info model.ExpireInfo, err error) {
	params := map[string]string{
		"user_id":        userId,
		"application_id": appId,
		"system_user_id": systemUserId,
	}
	_, err = s.authClient.Get(ValidateApplicationPermissionURL, &info, params)
	return
}

const actionConnect = "connect"

func (s *PAMService) ValidateAssetConnectPermission(userId, assetId, systemUserId string) (info model.ExpireInfo, err error) {
	params := map[string]string{
		"user_id":        userId,
		"asset_id":       assetId,
		"system_user_id": systemUserId,
		"action_name":    actionConnect,
	}
	_, err = s.authClient.Get(ValidateUserAssetPermissionURL, &info, params)
	return
}

func (s *PAMService) ValidateJoinSessionPermission(userId, sessionId string) (result model.ValidateResult, err error) {
	data := map[string]string{
		"user_id":    userId,
		"session_id": sessionId,
	}
	_, err = s.authClient.Post(JoinRoomValidateURL, data, &result)
	return
}
