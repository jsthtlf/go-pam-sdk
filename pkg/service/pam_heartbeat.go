package service

import (
	"github.com/jsthtlf/go-pam-sdk/pkg/common"
	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

func (s *PAMService) TerminalHeartBeat(sIds []string) (res []model.TerminalTask, err error) {
	data := model.HeartbeatData{
		SessionOnlineIds: sIds,
		CpuUsed:          common.CpuLoad1Usage(),
		MemoryUsed:       common.MemoryUsagePercent(),
		DiskUsed:         common.DiskUsagePercent(),
		SessionOnline:    len(sIds),
	}
	_, err = s.authClient.Post(TerminalHeartBeatURL, data, &res)
	return
}
