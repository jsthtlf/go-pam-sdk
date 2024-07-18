package confirm

import "github.com/jsthtlf/go-pam-sdk/pkg/model"

type confirmOptions struct {
	user       model.User
	systemUser model.SystemUserAuthInfo

	targetType string
	targetID   string
}

type Option func(*confirmOptions)

func WithUser(user model.User) Option {
	return func(option *confirmOptions) {
		option.user = user
	}
}

func WithSystemUser(sysUser model.SystemUserAuthInfo) Option {
	return func(option *confirmOptions) {
		option.systemUser = sysUser
	}
}

func WithTargetType(targetType string) Option {
	return func(option *confirmOptions) {
		option.targetType = targetType
	}
}

func WithTargetID(targetID string) Option {
	return func(option *confirmOptions) {
		option.targetID = targetID
	}
}
