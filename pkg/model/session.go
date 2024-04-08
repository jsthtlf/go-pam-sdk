package model

import (
	"github.com/jsthtlf/go-pam-sdk/pkg/utils"
)

type Session struct {
	ID           string        `json:"id"`
	User         string        `json:"user"` // "%s(%s)" Name Username
	Asset        string        `json:"asset"`
	SystemUser   string        `json:"system_user"`
	LoginFrom    string        `json:"login_from"`
	RemoteAddr   string        `json:"remote_addr"`
	Protocol     string        `json:"protocol"`
	DateStart    utils.UTCTime `json:"date_start"`
	OrgID        string        `json:"org_id"`
	UserID       string        `json:"user_id"`
	AssetID      string        `json:"asset_id"`
	SystemUserID string        `json:"system_user_id"`
}
