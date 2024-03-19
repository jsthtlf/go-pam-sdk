package service

import (
	"time"

	"github.com/jsthtlf/go-pam-sdk/pkg/httplib"
)

type option struct {
	// default http://127.0.0.1:8080
	CoreHost string
	TimeOut  time.Duration
	sign     httplib.AuthSign
}

type Option func(*option)

func PAMCoreHost(coreHost string) Option {
	return func(o *option) {
		o.CoreHost = coreHost
	}
}

func PAMTimeOut(t time.Duration) Option {
	return func(o *option) {
		o.TimeOut = t
	}
}

func PAMAccessKey(keyID, secretID string) Option {
	return func(o *option) {
		o.sign = &httplib.SigAuth{
			KeyID:    keyID,
			SecretID: secretID,
		}
	}
}
