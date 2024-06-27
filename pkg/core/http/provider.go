package http

import (
	"errors"
	"net/http"
	"time"

	"github.com/jsthtlf/go-pam-sdk/pkg/core"
	"github.com/jsthtlf/go-pam-sdk/pkg/httplib"
)

var _ core.Provider = (*httpProvider)(nil)

const (
	minTimeOut = time.Second * 30

	orgHeaderKey   = "X-PAM-ORG"
	orgHeaderValue = "ROOT"

	langCookieKey   = "django_language"
	langCookieValue = "en"
)

var defaultOptions = &options{
	Host:    "127.0.0.1",
	TimeOut: minTimeOut,
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
	httpClient.SetCookie(langCookieKey, langCookieValue)

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

func (p *httpProvider) CloneClient() httplib.Client {
	return p.authClient.Clone()
}

func (p *httpProvider) Copy() core.Provider {
	client := p.authClient.Clone()
	if p.opt.sign != nil {
		client.SetAuthSign(p.opt.sign)
	}
	client.SetCookie(langCookieKey, langCookieValue)
	client.SetHeader(orgHeaderKey, orgHeaderValue)

	return &httpProvider{
		authClient: &client,
		opt:        p.opt,
	}
}

func (p *httpProvider) SetCookie(name, value string) {
	p.authClient.SetCookie(name, value)
}

func (p *httpProvider) get(reqUrl string, res interface{}, params ...map[string]string) (resp *http.Response, err error) {
	resp, err = p.authClient.Get(reqUrl, res, params...)
	if p.needRequestAgain(err) {
		resp, err = p.authClient.Get(reqUrl, res, params...)
	}
	return
}

func (p *httpProvider) post(reqUrl string, data interface{}, res interface{}, params ...map[string]string) (resp *http.Response, err error) {
	resp, err = p.authClient.Post(reqUrl, data, res, params...)
	if p.needRequestAgain(err) {
		resp, err = p.authClient.Post(reqUrl, data, res, params...)
	}
	return
}

func (p *httpProvider) delete(reqUrl string, res interface{}, params ...map[string]string) (resp *http.Response, err error) {
	resp, err = p.authClient.Delete(reqUrl, res, params...)
	if p.needRequestAgain(err) {
		resp, err = p.authClient.Delete(reqUrl, res, params...)
	}
	return
}

func (p *httpProvider) put(reqUrl string, data interface{}, res interface{}, params ...map[string]string) (resp *http.Response, err error) {
	resp, err = p.authClient.Put(reqUrl, data, res, params...)
	if p.needRequestAgain(err) {
		resp, err = p.authClient.Put(reqUrl, data, res, params...)
	}
	return
}

func (p *httpProvider) patch(reqUrl string, data interface{}, res interface{}, params ...map[string]string) (resp *http.Response, err error) {
	resp, err = p.authClient.Patch(reqUrl, data, res, params...)
	if p.needRequestAgain(err) {
		resp, err = p.authClient.Patch(reqUrl, data, res, params...)
	}
	return
}

func (p *httpProvider) postFileWithFields(reqUrl string, gFile string, fields map[string]string, res interface{}) (err error) {
	err = p.authClient.PostFileWithFields(reqUrl, gFile, fields, res)
	if p.needRequestAgain(err) {
		err = p.authClient.PostFileWithFields(reqUrl, gFile, fields, res)
	}
	return err
}

func (p *httpProvider) needRequestAgain(err error) bool {
	if err == nil {
		return false
	}

	respErr := &httplib.ResponseError{}
	if errors.As(err, &respErr) {
		if respErr.HasCode(httplib.CodeAuthFailed) {
			err = p.signupAgain(err)
			if err == nil {
				return true
			}
		}
	}

	return false
}
