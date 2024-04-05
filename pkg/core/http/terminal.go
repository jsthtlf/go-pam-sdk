package http

import (
	"fmt"

	"github.com/jsthtlf/go-pam-sdk/pkg/common"
	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

func (p *httpProvider) GetTerminalConfig() (conf model.TerminalConfig, err error) {
	_, err = p.authClient.Get(UrlTerminalConfig, &conf)
	return
}

func (p *httpProvider) HeartBeat(sIds []string) (res []model.TerminalTask, err error) {
	data := model.HeartbeatData{
		SessionOnlineIds: sIds,
		CpuUsed:          common.CpuLoad1Usage(),
		MemoryUsed:       common.MemoryUsagePercent(),
		DiskUsed:         common.DiskUsagePercent(),
		SessionOnline:    len(sIds),
	}
	_, err = p.authClient.Post(UrlTerminalHeartBeat, data, &res)
	return
}

func (p *httpProvider) GetPublicSetting() (result model.PublicSetting, err error) {
	_, err = p.authClient.Get(UrlPublicSetting, &result)
	return
}

func (p *httpProvider) FinishTask(tid string) error {
	data := map[string]bool{"is_finished": true}
	Url := fmt.Sprintf(UrlTerminalTask, tid)
	_, err := p.authClient.Patch(Url, data, nil)
	return err
}
