package core

import (
	"github.com/jsthtlf/go-pam-sdk/pkg/common"
	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

type (
	Provider interface {
		UserProvider
		TerminalProvider
		AssetProvider
		ApplicationProvider
		SystemUserProvider
		AuditProvider
		SessionProvider
		PermissionProvider
		ShareRoomProvider
		TicketProvider
		TokenProvider

		Copy() Provider
		SetCookie(name, value string)
	}

	UserProvider interface {
		GetUserById(uid string) (user *model.User, err error)
		GetProfile() (user *model.User, err error)
		CheckUserCookie(cookies map[string]string) (user *model.User, err error)
		GetShareUserInfo(query string) (res []*model.MiniUser, err error)
	}

	TerminalProvider interface {
		Register() error
		GetTerminalConfig() (conf model.TerminalConfig, err error)
		HeartBeat(sIds []string) (res []model.TerminalTask, err error)
		GetPublicSetting() (result model.PublicSetting, err error)
		FinishTask(tid string) error
	}

	AssetProvider interface {
		GetAssetById(assetId string) (asset model.Asset, err error)
		GetAssetPlatform(assetId string) (platform model.Platform, err error)
		GetDomainGateways(domainId string) (domain model.Domain, err error)
	}

	ApplicationProvider interface {
		GetApplicationById(appId string) (app model.Application, err error)
		GetRemoteApplicationById(remoteAppId string) (remoteApp model.RemoteAPP, err error)
	}

	SystemUserProvider interface {
		GetSystemUserById(systemUserId string) (sysUser model.SystemUser, err error)
		GetUserApplicationAuthInfo(systemUserID, appID, userID, username string) (info model.SystemUserAuthInfo, err error)
		GetUserApplicationSystemUsers(userId, appId string) (res []model.SystemUser, err error)
		GetSystemUserAuthById(systemUserId, assetId, userId, username string) (info model.SystemUserAuthInfo, err error)
	}

	AuditProvider interface {
		CreateFileOperationLog(data model.FTPLog) (err error)
		CreateSessionCommand(commands []*model.Command) (err error)
		CreateNotifyCommand(commands []*model.Command) (err error)
		GetSystemUserFilterRules(systemUserID string) (rules []model.FilterRule, err error)
		GetCommandFilterRules(userId, sysId, assetId, appId string) (rules []model.FilterRule, err error)
	}

	SessionProvider interface {
		CreateSession(sess model.Session) error
		SessionSuccess(sid string) error
		SessionFailed(sid string, err error) error
		SessionDisconnect(sid string) error
		SessionFinished(sid string, time common.UTCTime) error
		GetSessionById(sid string) (data model.Session, err error)
		CreateSessionTicketRelation(sid, ticketId string) (err error)
		UploadReplay(sid, gZipFile string) error
		FinishReply(sid string) error
	}

	PermissionProvider interface {
		GetPermission(userId, assetId, systemUserId string) (perms model.Permission, err error)
		ValidateRemoteAppPermission(userId, remoteAppId, systemUserId string) (info model.ExpireInfo, err error)
		ValidateApplicationPermission(userId, appId, systemUserId string) (info model.ExpireInfo, err error)
		ValidateAssetConnectPermission(userId, assetId, systemUserId string) (info model.ExpireInfo, err error)
		ValidateJoinSessionPermission(userId, sessionId string) (result model.ValidateResult, err error)
		SearchPermAsset(userId, key string) (res model.AssetList, err error)
		GetSystemUsersByUserIdAndAssetId(userId, assetId string) (sysUsers []model.SystemUser, err error)
		GetAllUserPermsAssets(userId string) ([]map[string]interface{}, error)
		GetUserPermsAssets(userID string, params model.PaginationParam) (resp model.PaginationResponse, err error)
		RefreshUserAllPermsAssets(userId string) ([]map[string]interface{}, error)
		GetUserAssetByID(userId, assetId string) (assets []model.Asset, err error)
		GetUserPermAssetsByIP(userId, assetIP string) (assets []model.Asset, err error)
		GetAllUserPermK8s(userId string) ([]map[string]interface{}, error)
		GetUserPermsMySQL(userId string, param model.PaginationParam) (resp model.PaginationResponse, err error)
		GetUserPermsDatabase(userId string, param model.PaginationParam, dbTypes ...string) (resp model.PaginationResponse, err error)
		GetUserPermsK8s(userId string, param model.PaginationParam) (resp model.PaginationResponse, err error)
		GetUserNodeAssets(userID, nodeID string, params model.PaginationParam) (resp model.PaginationResponse, err error)
		GetUserNodes(userId string) (nodes model.NodeList, err error)
		RefreshUserNodes(userId string) (nodes model.NodeList, err error)
		GetNodeTreeByUserAndNodeKey(userID, nodeKey string) (nodeTrees model.NodeTreeList, err error)
	}

	ShareRoomProvider interface {
		CreateShareRoom(sessionId string, expired int, users []string) (res model.SharingSession, err error)
		JoinShareRoom(data model.SharePostData) (res model.ShareRecord, err error)
		FinishShareRoom(recordId string) (err error)
	}

	TicketProvider interface {
		SubmitCommandConfirm(sid string, ruleId string, cmd string) (res model.CommandTicketInfo, err error)
		CheckIfNeedAssetLoginConfirm(userId, assetId, systemUserId, sysUsername string) (res model.AssetLoginTicketInfo, err error)
		CheckIfNeedAppConnectionConfirm(userID, assetID, systemUserID string) (bool, error)
		CancelConfirmByRequestInfo(req model.ReqInfo) (err error)
		CheckConfirmStatusByRequestInfo(req model.ReqInfo) (res model.TicketState, err error)
	}

	TokenProvider interface {
		GetTokenAsset(token string) (tokenUser model.TokenUser, err error)
		GetConnectTokenAuth(token string) (resp model.ConnectTokenInfo, err error)
		RenewalToken(token string) (resp model.TokenRenewalResponse, err error)
	}
)
