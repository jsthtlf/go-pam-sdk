package http

import (
	"time"

	"github.com/jsthtlf/go-pam-sdk/pkg/httplib"
)

type options struct {
	Host           string
	TimeOut        time.Duration
	BootstrapToken string

	TerminalName    string
	TerminalComment string
	TerminalType    string
	AccessKeyPath   string

	sign httplib.AuthSign
}

type Option func(*options)

func WithHost(host string) Option {
	return func(o *options) {
		o.Host = host
	}
}

func WithTimeOut(t time.Duration) Option {
	return func(o *options) {
		o.TimeOut = t
	}
}

func WithBootstrapToken(token string) Option {
	return func(o *options) {
		o.BootstrapToken = token
	}
}

func WithTerminalName(name string) Option {
	return func(o *options) {
		o.TerminalName = name
	}
}

func WithTerminalComment(comment string) Option {
	return func(o *options) {
		o.TerminalComment = comment
	}
}

func WithTerminalType(typ string) Option {
	return func(o *options) {
		o.TerminalType = typ
	}
}

func WithAccessKeyPath(path string) Option {
	return func(o *options) {
		o.AccessKeyPath = path
	}
}
