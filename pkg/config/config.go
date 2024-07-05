package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jsthtlf/go-pam-sdk/pkg/logger"
	"github.com/jsthtlf/go-pam-sdk/pkg/utils"

	"github.com/spf13/viper"
)

const (
	TerminalDefault = "pam-default"
	TerminalDb      = "db"
	TerminalDbWeb   = "db-web-ssh"
	TerminalRdp     = "rdp"
	TerminalRdpWeb  = "rdp-web"
)

var config *Config

type Config struct {
	Name    string `mapstructure:"TERMINAL_NAME"`
	Comment string `mapstructure:"TERMINAL_COMMENT"`

	CoreHost       string `mapstructure:"CORE_HOST"`
	BootstrapToken string `mapstructure:"BOOTSTRAP_TOKEN"`
	LogLevel       string `mapstructure:"LOG_LEVEL"`

	TerminalType string

	RootPath          string
	DataFolderPath    string
	LogDirPath        string
	KeyFolderPath     string
	AccessKeyFilePath string
	ReplayFolderPath  string
}

func Initial() *Config {
	if config == nil {
		cfg := getDefaultConfig()
		loadConfigFromEnv(&cfg)
		config = &cfg

		logger.Initial(cfg.LogLevel, cfg.LogDirPath)
	}
	return config
}

func getDefaultConfig() Config {
	rootPath := getPwdDirPath()
	dataFolderPath := filepath.Join(rootPath, "data")
	replayFolderPath := filepath.Join(dataFolderPath, "replays")
	LogDirPath := filepath.Join(dataFolderPath, "logs")
	keyFolderPath := filepath.Join(dataFolderPath, "keys")
	accessKeyFilePath := filepath.Join(keyFolderPath, ".access_key")

	folders := []string{dataFolderPath, replayFolderPath, keyFolderPath, LogDirPath}
	for i := range folders {
		if err := utils.EnsureDirExist(folders[i]); err != nil {
			log.Fatalf("Create folder failed: %s", err)
		}
	}
	return Config{
		Name:    "",
		Comment: "Pam terminal",

		CoreHost:       "http://localhost:8080",
		BootstrapToken: "",
		LogLevel:       "INFO",

		TerminalType: TerminalDefault,

		AccessKeyFilePath: accessKeyFilePath,
		RootPath:          rootPath,
		DataFolderPath:    dataFolderPath,
		LogDirPath:        LogDirPath,
		KeyFolderPath:     keyFolderPath,
		ReplayFolderPath:  replayFolderPath,
	}
}

func getPwdDirPath() string {
	if rootPath, err := os.Getwd(); err == nil {
		return rootPath
	}
	return ""
}

func loadConfigFromEnv(conf *Config) {
	viper.AutomaticEnv()
	envViper := viper.New()
	for _, item := range os.Environ() {
		envItem := strings.SplitN(item, "=", 2)
		if len(envItem) == 2 {
			envViper.Set(envItem[0], viper.Get(envItem[0]))
		}
	}
	if err := envViper.Unmarshal(conf); err == nil {
		log.Println("Load config from env: success")
	}
}
