package confirm

import (
	"context"
	"time"

	"github.com/jsthtlf/go-pam-sdk/pkg/logger"
	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

type loginConfirmProvider interface {
	CheckIfNeedAssetLoginConfirm(userId, assetId, systemUserId, sysUsername string) (res model.AssetLoginTicketInfo, err error)
	CheckIfNeedAppConnectionConfirm(userID, assetID, systemUserID string) (bool, error)
	CancelConfirmByRequestInfo(req model.ReqInfo) (err error)
	CheckConfirmStatusByRequestInfo(req model.ReqInfo) (res model.TicketState, err error)
}

type Status int

const (
	StatusUnknown Status = iota
	StatusApprove
	StatusReject
	StatusCancel
)

var defaultOptions = &confirmOptions{
	attempts:     3,
	attemptDelay: 5 * time.Second,
}

func NewLoginConfirm(provider loginConfirmProvider, opts ...Option) LoginConfirmService {
	opt := defaultOptions

	for _, setter := range opts {
		setter(opt)
	}

	return LoginConfirmService{
		provider: provider,
		opts:     opt,
	}
}

type LoginConfirmService struct {
	provider loginConfirmProvider

	opts *confirmOptions

	checkReqInfo  model.ReqInfo
	cancelReqInfo model.ReqInfo

	reviewers       []string
	ticketDetailUrl string
	processor       string
	ticketId        string
	canceled        bool
}

func (c *LoginConfirmService) CheckIsNeedLoginConfirm() (bool, error) {
	userID := c.opts.user.ID
	systemUserID := c.opts.systemUser.ID
	systemUsername := c.opts.systemUser.Username
	targetID := c.opts.targetID
	switch c.opts.targetType {
	case model.AppType:
		return c.provider.CheckIfNeedAppConnectionConfirm(userID, targetID, systemUserID)
	default:
		res, err := c.provider.CheckIfNeedAssetLoginConfirm(userID, targetID,
			systemUserID, systemUsername)
		if err != nil {
			return false, err
		}
		c.ticketId = res.TicketId
		c.reviewers = res.Reviewers
		c.checkReqInfo = res.CheckReq
		c.cancelReqInfo = res.CloseReq
		c.ticketDetailUrl = res.TicketDetailUrl
		return res.NeedConfirm, nil
	}
}

func (c *LoginConfirmService) WaitLoginConfirm(ctx context.Context) Status {
	return c.waitConfirmFinish(ctx)
}

func (c *LoginConfirmService) GetReviewers() []string {
	reviewers := make([]string, len(c.reviewers))
	copy(reviewers, c.reviewers)
	return reviewers
}

func (c *LoginConfirmService) GetTicketUrl() string {
	return c.ticketDetailUrl
}

func (c *LoginConfirmService) GetProcessor() string {
	return c.processor
}

func (c *LoginConfirmService) GetTicketId() string {
	return c.ticketId
}

func (c *LoginConfirmService) waitConfirmFinish(ctx context.Context) Status {
	t := time.NewTicker(c.opts.attemptDelay)
	attemp := 0
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			c.cancelConfirm()
			return StatusCancel
		case <-t.C:
			statusRes, err := c.provider.CheckConfirmStatusByRequestInfo(c.checkReqInfo)
			if err != nil {
				logger.Errorf("Check confirm status failed: %s", err)
				if attemp > c.opts.attempts {
					return StatusUnknown
				}
				attemp++
				continue
			}
			switch statusRes.State {
			case model.TicketOpen:
				continue
			case model.TicketApproved:
				c.processor = statusRes.Processor
				return StatusApprove
			case model.TicketRejected:
				c.processor = statusRes.Processor
				return StatusReject
			case model.TicketCanceled:
				return StatusCancel
			default:
				logger.Errorf("Receive unknown login confirm status: %s", statusRes.Status)
				return StatusUnknown
			}
		}
	}
}

func (c *LoginConfirmService) cancelConfirm() {
	if c.canceled {
		return
	}
	if err := c.provider.CancelConfirmByRequestInfo(c.cancelReqInfo); err != nil {
		logger.Errorf("Cancel confirm request failed: %s", err)
	}
	c.canceled = true
}

func (c *LoginConfirmService) CancelConfirm() {
	c.cancelConfirm()
}
