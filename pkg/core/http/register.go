package http

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jsthtlf/go-pam-sdk/pkg/httplib"
	"github.com/jsthtlf/go-pam-sdk/pkg/logger"
	"github.com/jsthtlf/go-pam-sdk/pkg/model"
)

func (p *httpProvider) Register() error {
	var key model.AccessKey
	attempts := 10

	if err := key.LoadFromFile(p.opt.AccessKeyPath); err != nil {
		logger.Error("Load access key failed: %v", err)
		return p.register(attempts)
	}

	return p.validAccesKey(attempts, key)
}

func (p *httpProvider) register(attempts int) error {
	for i := 0; i < attempts; i++ {
		terminal, err := p.registerAccount()
		if err != nil {
			logger.Error(err)
			time.Sleep(time.Second * 3)
			continue
		}

		p.setSign(terminal.ServiceAccount.AccessKey.ID, terminal.ServiceAccount.AccessKey.Secret)

		if err := terminal.ServiceAccount.AccessKey.SaveToFile(p.opt.AccessKeyPath); err != nil {
			logger.Error("Error while save access key: %v", err)
		}

		return nil
	}

	return errors.New("attempts register account exceeded")
}

func (p *httpProvider) registerAccount() (res model.Terminal, err error) {
	regClient := p.authClient.Clone()
	regClient.SetHeader("Authorization", fmt.Sprintf("BootstrapToken %s", p.opt.BootstrapToken))
	data := map[string]string{
		"name":    p.opt.TerminalName,
		"comment": p.opt.TerminalComment,
		"type":    p.opt.TerminalType}
	_, err = regClient.Post(UrlTerminalRegister, data, &res)
	return
}

func (p *httpProvider) validAccesKey(attempts int, key model.AccessKey) error {
	for i := 0; i < attempts; i++ {
		if err := validAccessKey(p.opt.Host, key); err != nil {
			switch {
			case errors.Is(err, ErrUnauthorized):
				logger.Error("Access key unauthorized, try to register new access key")
				return p.register(attempts)
			default:
				logger.Error("Check access key failed: %v", err)
			}
			time.Sleep(time.Second * 3)
			continue
		}

		p.setSign(key.ID, key.Secret)
		return nil
	}

	return errors.New("attempts valid access key exceeded")
}

func (p *httpProvider) setSign(keyID, secretID string) {
	p.opt.sign = &httplib.SigAuth{
		KeyID:    keyID,
		SecretID: secretID,
	}

	p.authClient.SetAuthSign(p.opt.sign)
}

func validAccessKey(host string, key model.AccessKey) error {
	client, err := httplib.NewClient(host, time.Second*30)
	if err != nil {
		return err
	}
	client.SetAuthSign(&httplib.SigAuth{
		KeyID:    key.ID,
		SecretID: key.Secret,
	})
	var (
		user model.User
		res  *http.Response
	)
	res, err = client.Get(UrlUserProfile, &user)
	if err != nil {
		if res == nil {
			return fmt.Errorf("%w: %v", ErrConnect, err)
		}
		if res.StatusCode == http.StatusUnauthorized {
			return ErrUnauthorized
		}
		return fmt.Errorf("%w: %v", ErrInvalid, err)
	}

	if user.ID == "" {
		return ErrInvalid
	}

	return nil
}
