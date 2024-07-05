package heartbeat

import (
	"time"

	"github.com/jsthtlf/go-pam-sdk/pkg/logger"
	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

type heartbeatProvider interface {
	HeartBeat(sIds []string) (res model.HeartbeatResponse, err error)
}

type sessionsCallback func() []string
type tasksCallback func(tasks []model.TerminalTask)

func Start(getSessions sessionsCallback, executeTasks tasksCallback, p heartbeatProvider) {
	interval := 5
	for {
		logger.Debug("Send terminal heartbeat...")
		data := getSessions()
		resp, err := p.HeartBeat(data)
		if err != nil {
			logger.Warn("Send terminal heartbeat failed: ", err)
			continue
		}
		interval = resp.NextHeartbeat
		if len(resp.Tasks) != 0 {
			go executeTasks(resp.Tasks)
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
