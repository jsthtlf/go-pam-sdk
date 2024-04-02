package recorder

import (
	"sync"
	"time"

	"github.com/jsthtlf/go-pam-sdk/pkg/logger"
	"github.com/jsthtlf/go-pam-sdk/pkg/model"
	"github.com/jsthtlf/go-pam-sdk/pkg/service"
	"github.com/jsthtlf/go-pam-sdk/pkg/storage"
)

type CommandRecorder struct {
	sessionID  string
	pamService *service.PAMService
	storage    storage.CommandStorage

	queue chan *model.Command
	done  chan struct{}
	wg    sync.WaitGroup
}

func NewCommandRecorder(sid string,
	pamService *service.PAMService,
	storage storage.CommandStorage) (*CommandRecorder, error) {
	return &CommandRecorder{
		sessionID:  sid,
		storage:    storage,
		queue:      make(chan *model.Command, 10),
		done:       make(chan struct{}),
		pamService: pamService,
	}, nil
}

func (c *CommandRecorder) RecordCommand(command *model.Command) {
	c.wg.Add(1)
	c.queue <- command
}

func (c *CommandRecorder) End() {
	select {
	case <-c.done:
		return
	default:
	}
	c.wg.Wait()

	close(c.done)
}

func (c *CommandRecorder) Record() {
	cmdList := make([]*model.Command, 0, 10)
	notificationList := make([]*model.Command, 0, 10)
	maxRetry := 0
	tick := time.NewTicker(time.Second * 3)
	defer tick.Stop()
	for {
		select {
		case <-c.done:
			if len(cmdList) == 0 {
				return
			}
		case p, ok := <-c.queue:
			if !ok {
				return
			}
			if p.RiskLevel == model.DangerLevel {
				notificationList = append(notificationList, p)
			}
			cmdList = append(cmdList, p)
			if len(cmdList) < 5 {
				continue
			}
		case <-tick.C:
			if len(cmdList) == 0 {
				continue
			}
		}
		if len(notificationList) > 0 {
			if err := c.pamService.NotifyCommand(notificationList); err == nil {
				notificationList = notificationList[:0]
			} else {
				logger.Errorf("Create notify command for session (%s) failed: %+v", c.sessionID, err)
			}
		}
		err := c.storage.BulkSave(cmdList)
		if err == nil {
			c.wg.Add(-len(cmdList))
			cmdList = cmdList[:0]
			maxRetry = 0
			continue
		}
		if err != nil {
			logger.Errorf("Bulk save command for session (%s) failed: %+v", c.sessionID, err)
		}

		if maxRetry > 5 {
			cmdList = cmdList[1:]
		}
		maxRetry++
	}
}
