package http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

func (p *httpProvider) SubmitCommandConfirm(sid string, ruleId string, cmd string) (res model.CommandTicketInfo, err error) {
	data := map[string]string{
		"session_id":         sid,
		"cmd_filter_rule_id": ruleId,
		"run_command":        cmd,
	}
	_, err = p.post(UrlCommandConfirm, data, &res)
	return
}

func (p *httpProvider) CheckIfNeedAssetLoginConfirm(userId, assetId, systemUserId, sysUsername string) (res model.AssetLoginTicketInfo, err error) {
	data := map[string]string{
		"user_id":              userId,
		"asset_id":             assetId,
		"system_user_id":       systemUserId,
		"system_user_username": sysUsername,
	}

	_, err = p.post(UrlAssetLoginConfirm, data, &res)
	return
}

func (p *httpProvider) CheckIfNeedAppConnectionConfirm(userID, assetID, systemUserID string) (bool, error) {

	return false, nil
}

func (p *httpProvider) CancelConfirmByRequestInfo(req model.ReqInfo) (err error) {
	res := make(map[string]interface{})
	err = p.sendRequestByRequestInfo(req, &res)
	return
}

func (p *httpProvider) CheckConfirmStatusByRequestInfo(req model.ReqInfo) (res model.TicketState, err error) {
	err = p.sendRequestByRequestInfo(req, &res)
	return
}

func (p *httpProvider) sendRequestByRequestInfo(req model.ReqInfo, res interface{}) (err error) {
	switch strings.ToUpper(req.Method) {
	case http.MethodGet:
		_, err = p.get(req.URL, res)
	case http.MethodDelete:
		_, err = p.delete(req.URL, res)
	default:
		err = fmt.Errorf("unsupport method %s", req.Method)
	}
	return
}
