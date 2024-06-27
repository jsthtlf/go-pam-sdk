package http

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

func (p *httpProvider) GetPermission(userId, assetId, systemUserId string) (perms model.Permission, err error) {
	params := map[string]string{
		"user_id":        userId,
		"asset_id":       assetId,
		"system_user_id": systemUserId,
	}
	_, err = p.get(UrlAssetPermsDetail, &perms, params)
	return
}

func (p *httpProvider) ValidateRemoteAppPermission(userId, remoteAppId, systemUserId string) (info model.ExpireInfo, err error) {
	return p.ValidateApplicationPermission(userId, remoteAppId, systemUserId)
}

func (p *httpProvider) ValidateApplicationPermission(userId, appId, systemUserId string) (info model.ExpireInfo, err error) {
	params := map[string]string{
		"user_id":        userId,
		"application_id": appId,
		"system_user_id": systemUserId,
	}
	_, err = p.get(UrlValidateAppPerms, &info, params)
	return
}

const actionConnect = "connect"

func (p *httpProvider) ValidateAssetConnectPermission(userId, assetId, systemUserId string) (info model.ExpireInfo, err error) {
	params := map[string]string{
		"user_id":        userId,
		"asset_id":       assetId,
		"system_user_id": systemUserId,
		"action_name":    actionConnect,
	}
	_, err = p.get(UrlValidateAssetPerms, &info, params)
	return
}

func (p *httpProvider) ValidateJoinSessionPermission(userId, sessionId string) (result model.ValidateResult, err error) {
	data := map[string]string{
		"user_id":    userId,
		"session_id": sessionId,
	}
	_, err = p.post(UrlValidateJoinRoom, data, &result)
	return
}

func (p *httpProvider) SearchPermAsset(userId, key string) (res model.AssetList, err error) {
	Url := fmt.Sprintf(UrlUserPermsAssets, userId)
	payload := map[string]string{"search": key}
	_, err = p.get(Url, &res, payload)
	return
}

func (p *httpProvider) GetSystemUsersByUserIdAndAssetId(userId, assetId string) (sysUsers []model.SystemUser, err error) {
	Url := fmt.Sprintf(UrlUserPermsAssetSystemUsers, userId, assetId)
	_, err = p.get(Url, &sysUsers)
	return
}

func (p *httpProvider) GetAllUserPermsAssets(userId string) ([]map[string]interface{}, error) {
	var params model.PaginationParam
	res, err := p.GetUserPermsAssets(userId, params)
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}

func (p *httpProvider) GetUserPermsAssets(userID string, params model.PaginationParam) (resp model.PaginationResponse, err error) {
	Url := fmt.Sprintf(UrlUserPermsAssets, userID)
	return p.getPaginationResult(Url, params)
}

func (p *httpProvider) RefreshUserAllPermsAssets(userId string) ([]map[string]interface{}, error) {
	var params model.PaginationParam
	params.Refresh = true
	res, err := p.GetUserPermsAssets(userId, params)
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}

func (p *httpProvider) GetUserAssetByID(userId, assetId string) (assets []model.Asset, err error) {
	params := map[string]string{
		"id": assetId,
	}
	Url := fmt.Sprintf(UrlUserPermsAssets, userId)
	_, err = p.get(Url, &assets, params)
	return
}

func (p *httpProvider) GetUserPermAssetsByIP(userId, assetIP string) (assets []model.Asset, err error) {
	params := map[string]string{
		"ip": assetIP,
	}
	reqUrl := fmt.Sprintf(UrlUserPermsAssets, userId)
	_, err = p.get(reqUrl, &assets, params)
	return
}

func (p *httpProvider) getPaginationResult(reqUrl string, param model.PaginationParam) (resp model.PaginationResponse, err error) {
	if param.PageSize < 0 {
		param.PageSize = 0
	}
	paramsArray := make([]map[string]string, 0, len(param.Searches)+2)
	for i := 0; i < len(param.Searches); i++ {
		paramsArray = append(paramsArray, map[string]string{
			"search": strings.TrimSpace(param.Searches[i]),
		})
	}

	params := map[string]string{
		"limit":  strconv.Itoa(param.PageSize),
		"offset": strconv.Itoa(param.Offset),
	}
	if param.Refresh {
		params["rebuild_tree"] = "1"
	}
	paramsArray = append(paramsArray, params)
	if param.PageSize > 0 {
		_, err = p.get(reqUrl, &resp, paramsArray...)
	} else {
		var data []map[string]interface{}
		_, err = p.get(reqUrl, &data, paramsArray...)
		resp.Data = data
		resp.Total = len(data)
	}
	return
}

func (p *httpProvider) GetAllUserPermK8s(userId string) ([]map[string]interface{}, error) {
	var param model.PaginationParam
	res, err := p.GetUserPermsK8s(userId, param)
	if err != nil {
		return nil, err
	}
	return res.Data, err
}

func (p *httpProvider) GetUserPermsMySQL(userId string, param model.PaginationParam) (resp model.PaginationResponse, err error) {
	reqUrl := fmt.Sprintf(UrlUserPermsApplications, userId, model.AppTypeMySQL)
	return p.getPaginationResult(reqUrl, param)
}

func (p *httpProvider) GetUserPermsDatabase(userId string, param model.PaginationParam, dbTypes ...string) (resp model.PaginationResponse, err error) {
	reqUrl := fmt.Sprintf(UrlUserPermsDatabase, userId, strings.Join(dbTypes, ","))
	return p.getPaginationResult(reqUrl, param)
}

func (p *httpProvider) GetUserPermsK8s(userId string, param model.PaginationParam) (resp model.PaginationResponse, err error) {
	reqUrl := fmt.Sprintf(UrlUserPermsApplications, userId, model.AppTypeK8s)
	return p.getPaginationResult(reqUrl, param)
}

func (p *httpProvider) GetUserNodeAssets(userID, nodeID string, params model.PaginationParam) (resp model.PaginationResponse, err error) {
	Url := fmt.Sprintf(UrlUserPermsNodeAssetsList, userID, nodeID)
	return p.getPaginationResult(Url, params)
}

func (p *httpProvider) GetUserNodes(userId string) (nodes model.NodeList, err error) {
	Url := fmt.Sprintf(UrlUserPermsNodesList, userId)
	_, err = p.get(Url, &nodes)
	return
}

func (p *httpProvider) RefreshUserNodes(userId string) (nodes model.NodeList, err error) {
	params := map[string]string{
		"rebuild_tree": "1",
	}
	Url := fmt.Sprintf(UrlUserPermsNodesList, userId)
	_, err = p.get(Url, &nodes, params)
	return
}

func (p *httpProvider) GetNodeTreeByUserAndNodeKey(userID, nodeKey string) (nodeTrees model.NodeTreeList, err error) {
	payload := map[string]string{}
	if nodeKey != "" {
		payload["key"] = nodeKey
	}
	apiURL := fmt.Sprintf(UrlUserPermsNodeTreeWithAsset, userID)
	_, err = p.get(apiURL, &nodeTrees, payload)
	return
}
