package http

import (
	"fmt"

	"github.com/jsthtlf/go-pam-sdk/pkg/model"
	"github.com/jsthtlf/go-pam-sdk/pkg/utils"
)

func (p *httpProvider) GetTerminalConfig() (conf model.TerminalConfig, err error) {
	_, err = p.get(UrlTerminalConfig, &conf)
	return
}

func (p *httpProvider) HeartBeat(sIds []string) (res model.HeartbeatResponse, err error) {
	data := model.HeartbeatData{
		SessionOnlineIds: sIds,
		CpuUsed:          utils.CpuLoad1Usage(),
		MemoryUsed:       utils.MemoryUsagePercent(),
		DiskUsed:         utils.DiskUsagePercent(),
		SessionOnline:    len(sIds),
	}
	_, err = p.post(UrlTerminalHeartBeat, data, &res)
	return
}

func (p *httpProvider) GetPublicSetting() (result model.PublicSetting, err error) {
	_, err = p.get(UrlPublicSetting, &result)
	return
}

func (p *httpProvider) FinishTask(tid string) error {
	data := map[string]bool{"is_finished": true}
	Url := fmt.Sprintf(UrlTerminalTask, tid)
	_, err := p.patch(Url, data, nil)
	return err
}

func (p *httpProvider) GetTerminalTasks() (tasks model.TerminalTasks, err error) {
	_, err = p.post(UrlTerminalTasks, nil, &tasks)
	return
}
