package http

import (
	"fmt"

	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

func (p *httpProvider) CreateFileOperationLog(data model.FTPLog) (err error) {
	_, err = p.authClient.Post(UrlFtpLogList, data, nil)
	return
}

func (p *httpProvider) CreateSessionCommand(commands []*model.Command) (err error) {
	_, err = p.authClient.Post(UrlSessionCommand, commands, nil)
	return
}

func (p *httpProvider) CreateNotifyCommand(commands []*model.Command) (err error) {
	_, err = p.authClient.Post(UrlSessionNotifyCommand, commands, nil)
	return
}

func (p *httpProvider) GetSystemUserFilterRules(systemUserID string) (rules []model.FilterRule, err error) {
	Url := fmt.Sprintf(UrlSystemUserCmdFilterRulesList, systemUserID)
	_, err = p.authClient.Get(Url, &rules)
	return
}

func (p *httpProvider) GetCommandFilterRules(userId, sysId, assetId, appId string) (rules []model.FilterRule, err error) {
	param := make(map[string]string)
	if userId != "" {
		param["user_id"] = userId
	}
	if sysId != "" {
		param["system_user_id"] = sysId
	}
	if assetId != "" {
		param["asset_id"] = assetId
	}
	if appId != "" {
		param["application_id"] = appId
	}

	_, err = p.authClient.Get(UrlCmdFilterRulesList, &rules, param)
	return
}
