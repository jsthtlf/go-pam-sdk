package service

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jsthtlf/go-pam-sdk/httplib"
	"github.com/jsthtlf/go-pam-sdk/model"
	//"github.com/davecgh/go-spew/spew"
)

func RegisterTerminalAccount(coreHost, token, name, comment, typeTerminal string) (res model.Terminal, err error) {
	client, err := httplib.NewClient(coreHost, time.Second*30)
	if err != nil {
		return model.Terminal{}, err
	}
	client.SetHeader("Authorization", fmt.Sprintf("BootstrapToken %s", token))
	data := map[string]string{"name": name,
		"comment": comment,
		"type":    typeTerminal}
	_, err = client.Post(TerminalRegisterURL, data, &res)
	return
}

func ValidAccessKey(coreHost string, key model.AccessKey) error {
	client, err := httplib.NewClient(coreHost, time.Second*30)
	if err != nil {
		return err
	}
	sign := httplib.SigAuth{
		KeyID:    key.ID,
		SecretID: key.Secret,
	}
	client.SetAuthSign(&sign)
	var (
		user model.User
		res  *http.Response
	)
	res, err = client.Get(UserProfileURL, &user)
	if err != nil {
		if res == nil {
			return fmt.Errorf("%w:%s", ErrConnect, err.Error())
		}
		if res.StatusCode == http.StatusUnauthorized {
			return ErrUnauthorized
		}
		return fmt.Errorf("%w: %s", ErrInvalid, err.Error())
	}
	if user.ID == "" {
		return ErrInvalid
	}
	return nil
}

var (
	ErrConnect      = errors.New("connect failed")
	ErrUnauthorized = errors.New("unauthorized")
	ErrInvalid      = errors.New("invalid user")
)
