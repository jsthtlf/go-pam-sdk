package tasker

import (
	"github.com/jsthtlf/go-pam-sdk/pkg/logger"
	"github.com/jsthtlf/go-pam-sdk/pkg/model"
	"time"
)

type taskProvider interface {
	GetTerminalTasks() (tasksList model.TerminalTasks, err error)
}

type tasksCallback func(tasks []model.TerminalTask)

func Start(executeTasks tasksCallback, p taskProvider) {
	interval := 5
	for {
		time.Sleep(time.Second * time.Duration(interval))
		logger.Debug("Get terminal tasks...")
		resp, err := p.GetTerminalTasks()
		if err != nil {
			logger.Warn("Get terminal tasks failed: ", err)
			continue
		}
		interval = resp.NextRequest
		if len(resp.Tasks) != 0 {
			go executeTasks(resp.Tasks)
		}
	}
}
