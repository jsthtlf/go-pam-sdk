package service

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/jsthtlf/go-pam-sdk/httplib"
	"github.com/jsthtlf/go-pam-sdk/model"
)

var AccessKeyUnauthorized = errors.New("access key unauthorized")

var ConnectErr = errors.New("api connect err")

const (
	minTimeOut = time.Second * 30

	orgHeaderKey   = "X-JMS-ORG"
	orgHeaderValue = "ROOT"
)

func NewAuthPAMService(opts ...Option) (*PAMService, error) {
	opt := option{
		CoreHost: "http://127.0.0.1:8080",
		TimeOut:  time.Minute,
	}
	for _, setter := range opts {
		setter(&opt)
	}
	if opt.TimeOut < minTimeOut {
		opt.TimeOut = minTimeOut
	}
	httpClient, err := httplib.NewClient(opt.CoreHost, opt.TimeOut)
	if err != nil {
		return nil, err
	}
	if opt.sign != nil {
		httpClient.SetAuthSign(opt.sign)
	}
	httpClient.SetHeader(orgHeaderKey, orgHeaderValue)
	return &PAMService{authClient: httpClient, opt: &opt}, nil
}

type PAMService struct {
	authClient *httplib.Client
	opt        *option

	sync.Mutex
}

func (s *PAMService) GetUserById(userID string) (user *model.User, err error) {
	url := fmt.Sprintf(UserDetailURL, userID)
	_, err = s.authClient.Get(url, &user)
	return
}

func (s *PAMService) GetProfile() (user *model.User, err error) {
	var res *http.Response
	res, err = s.authClient.Get(UserProfileURL, &user)
	if res == nil && err != nil {
		return nil, fmt.Errorf("%w:%s", ConnectErr, err.Error())
	}
	if res != nil && res.StatusCode == http.StatusUnauthorized {
		return user, AccessKeyUnauthorized
	}
	return user, err
}

func (s *PAMService) GetTerminalConfig() (conf model.TerminalConfig, err error) {
	_, err = s.authClient.Get(TerminalConfigURL, &conf)
	return
}

func (s *PAMService) CloneClient() httplib.Client {
	return s.authClient.Clone()
}

func (s *PAMService) Copy() *PAMService {
	client := s.authClient.Clone()
	if s.opt.sign != nil {
		client.SetAuthSign(s.opt.sign)
	}
	client.SetHeader(orgHeaderKey, orgHeaderValue)
	return &PAMService{
		authClient: &client,
		opt:        s.opt,
	}
}

func (s *PAMService) SetCookie(name, value string) {
	s.authClient.SetCookie(name, value)
}
