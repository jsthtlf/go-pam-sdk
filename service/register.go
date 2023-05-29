package service

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jsthtlf/go-pam-sdk/config"
	"github.com/jsthtlf/go-pam-sdk/httplib"
	"github.com/jsthtlf/go-pam-sdk/logger"
	"github.com/jsthtlf/go-pam-sdk/model"
)

func MustPAMService() *PAMService {
	key := MustLoadValidAccessKey()
	pamService, err := NewAuthPAMService(PAMCoreHost(
		config.GetCurrentConfig().CoreHost), PAMTimeOut(30*time.Second),
		PAMAccessKey(key.ID, key.Secret),
	)
	if err != nil {
		logger.Fatal("Error while creating terminal: %s" + err.Error())
		os.Exit(1)
	}
	return pamService
}

func MustLoadValidAccessKey() model.AccessKey {
	conf := config.GetCurrentConfig()
	var key model.AccessKey
	if err := key.LoadFromFile(conf.AccessKeyFilePath); err != nil {
		return MustRegisterTerminalAccount()
	}
	// accessKey
	return MustValidKey(key)
}

func MustRegisterTerminalAccount() (key model.AccessKey) {
	conf := config.GetCurrentConfig()
	for i := 0; i < 10; i++ {
		terminal, err := RegisterTerminalAccount(conf.CoreHost,
			conf.Name, conf.BootstrapToken, conf.Comment, conf.TerminalType)
		if err != nil {
			logger.Error(err.Error())
			time.Sleep(5 * time.Second)
			continue
		}
		key.ID = terminal.ServiceAccount.AccessKey.ID
		key.Secret = terminal.ServiceAccount.AccessKey.Secret
		if err := key.SaveToFile(conf.AccessKeyFilePath); err != nil {
			logger.Error("Error while save access key: " + err.Error())
		}
		return key
	}
	logger.Error("Error while registration terminal")
	os.Exit(1)
	return
}

func MustValidKey(key model.AccessKey) model.AccessKey {
	conf := config.GetCurrentConfig()
	for i := 0; i < 10; i++ {
		if err := ValidAccessKey(conf.CoreHost, key); err != nil {
			switch {
			case errors.Is(err, ErrUnauthorized):
				logger.Error("Access key unauthorized, try to register new access key")
				return MustRegisterTerminalAccount()
			default:
				logger.Error("Check access key failed: " + err.Error())
			}
			time.Sleep(5 * time.Second)
			continue
		}
		return key
	}
	logger.Error("Error while checking access key")
	os.Exit(1)
	return key
}

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
