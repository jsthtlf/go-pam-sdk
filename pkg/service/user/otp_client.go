package user

import (
	"github.com/jsthtlf/go-pam-sdk/pkg/httplib"
	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

const (
	UrlUserAuthToken     = "/api/v1/authentication/tokens/"
	UrlUserAuthConfirm   = "/api/v1/authentication/login-confirm-ticket/status/"
	UrlUserAuthMFASelect = "/api/v1/authentication/mfa/select/"
)

var defaultOptions = &options{}

func NewOTPClient(setters ...Option) *OTPClient {
	opts := defaultOptions
	for _, setter := range setters {
		setter(opts)
	}
	if opts.RemoteAddr != "" {
		opts.client.SetHeader("X-Forwarded-For", opts.RemoteAddr)
	}
	if opts.LoginType != "" {
		opts.client.SetHeader("X-PAM-LOGIN-TYPE", opts.LoginType)
	}
	return &OTPClient{
		client: opts.client,
		Opts:   opts,
	}
}

type OTPClient struct {
	client *httplib.Client
	Opts   *options
}

func (c *OTPClient) SetOption(setters ...Option) {
	for _, setter := range setters {
		setter(c.Opts)
	}
}

func (c *OTPClient) GetAPIToken() (resp AuthResponse, err error) {
	data := map[string]string{
		"username":    c.Opts.Username,
		"password":    c.Opts.Password,
		"public_key":  c.Opts.PublicKey,
		"remote_addr": c.Opts.RemoteAddr,
		"login_type":  c.Opts.LoginType,
	}
	_, err = c.client.Post(UrlUserAuthToken, data, &resp)
	return
}

func (c *OTPClient) CheckConfirmAuthStatus() (resp AuthResponse, err error) {
	_, err = c.client.Get(UrlUserAuthConfirm, &resp)
	return
}

func (c *OTPClient) CancelConfirmAuth() (err error) {
	_, err = c.client.Delete(UrlUserAuthConfirm, nil)
	return
}

func (c *OTPClient) SendOTPRequest(optReq *OTPRequest) (resp AuthResponse, err error) {
	_, err = c.client.Post(optReq.ReqURL, optReq.ReqBody, &resp)
	return
}

func (c *OTPClient) SelectMFAChoice(mfaType string) (err error) {
	data := map[string]string{
		"type": mfaType,
	}
	_, err = c.client.Post(UrlUserAuthMFASelect, data, nil)
	return
}

type OTPRequest struct {
	ReqURL  string
	ReqBody map[string]interface{}
}

type DataResponse struct {
	Choices []string `json:"choices,omitempty"`
	Url     string   `json:"url,omitempty"`
}

type AuthResponse struct {
	Err  string       `json:"error,omitempty"`
	Msg  string       `json:"msg,omitempty"`
	Data DataResponse `json:"data,omitempty"`

	Username    string `json:"username,omitempty"`
	Token       string `json:"token,omitempty"`
	Keyword     string `json:"keyword,omitempty"`
	DateExpired string `json:"date_expired,omitempty"`

	User model.User `json:"user,omitempty"`
}
