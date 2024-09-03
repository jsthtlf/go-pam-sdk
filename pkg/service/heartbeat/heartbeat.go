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

func Start(getSessions sessionsCallback, p heartbeatProvider) {
	interval := 5
	for {
		time.Sleep(time.Second * time.Duration(interval))
		logger.Debug("Send terminal heartbeat...")
		data := getSessions()
		resp, err := p.HeartBeat(data)
		if err != nil {
			logger.Warn("Send terminal heartbeat failed: ", err)
			continue
		}
		interval = resp.NextHeartbeat
	}
}
