package getconfig

import (
	"time"

	"github.com/jsthtlf/go-pam-sdk/pkg/logger"
	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

type configProvider interface {
	GetTerminalConfig() (conf model.TerminalConfig, err error)
}

type updateConfigCallback func(config model.TerminalConfig)

func Start(interval int, updateConfig updateConfigCallback, p configProvider) {
	for {
		logger.Debug("Updating terminal config...")
		conf, err := p.GetTerminalConfig()
		if err != nil {
			logger.Errorf("Update terminal config failed: %s", err)
			continue
		}

		updateConfig(conf)
		time.Sleep(time.Second * time.Duration(interval))
	}
}
