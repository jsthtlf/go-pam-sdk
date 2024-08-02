package main

import (
	"github.com/jsthtlf/go-pam-sdk/pkg/config"
	"github.com/jsthtlf/go-pam-sdk/pkg/core/http"
	"github.com/jsthtlf/go-pam-sdk/pkg/logger"
)

func main() {
	cfg := config.Initial()

	logger.Debug("Terminal is connecting...")
	provider, err := http.New(
		http.WithHost(cfg.CoreHost),
		http.WithBootstrapToken(cfg.BootstrapToken),
		http.WithTerminalName(cfg.Name),
		http.WithTerminalComment(cfg.Comment),
		http.WithTerminalType(cfg.TerminalType),
		http.WithAccessKeyPath(cfg.AccessKeyFilePath))
	if err != nil {
		logger.Fatal("Create terminal provider failed: ", err)
	}
	if err = provider.Register(); err != nil {
		logger.Fatal("Register terminal failed: ", err)
	}
}
