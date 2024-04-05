package http

import (
	"time"

	"github.com/jsthtlf/go-pam-sdk/pkg/core"
	"github.com/jsthtlf/go-pam-sdk/pkg/httplib"
)

var _ core.Provider = (*httpProvider)(nil)

const (
	minTimeOut = time.Second * 30

	orgHeaderKey   = "X-PAM-ORG"
	orgHeaderValue = "ROOT"
)

var defaultOptions = &options{
	Host:            "127.0.0.1",
	TimeOut:         minTimeOut,
	BootstrapToken:  "",
	TerminalName:    "asd", // TODO
	TerminalComment: "",
	TerminalType:    "pam-default", // TODO list
	AccessKeyPath:   "",
}

func New(opts ...Option) (core.Provider, error) {
	opt := defaultOptions

	for _, setter := range opts {
		setter(opt)
	}

	if opt.TimeOut < minTimeOut {
		opt.TimeOut = minTimeOut
	}

	httpClient, err := httplib.NewClient(opt.Host, opt.TimeOut)
	if err != nil {
		return nil, err
	}

	httpClient.SetHeader(orgHeaderKey, orgHeaderValue)

	p := &httpProvider{
		authClient: httpClient,
		opt:        opt,
	}

	return p, nil
}

type httpProvider struct {
	authClient *httplib.Client
	opt        *options
}

func (p *httpProvider) Copy() core.Provider {
	client := p.authClient.Clone()
	if p.opt.sign != nil {
		client.SetAuthSign(p.opt.sign)
	}
	client.SetHeader(orgHeaderKey, orgHeaderValue)

	return &httpProvider{
		authClient: &client,
		opt:        p.opt,
	}
}

func (p *httpProvider) SetCookie(name, value string) {
	p.authClient.SetCookie(name, value)
}
