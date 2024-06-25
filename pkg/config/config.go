package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jsthtlf/go-pam-sdk/pkg/utils"

	"github.com/spf13/viper"
)

const (
	hostEnvKey = "SERVER_HOSTNAME"

	defaultNameMaxLen = 128

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
	}
	return config
}

func getDefaultConfig() Config {
	defaultName := getDefaultName()

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
		Name:    defaultName,
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

/*
SERVER_HOSTNAME: Имя переменной окружения, может использоваться для настройки префикса зарегистрированного имени по умолчанию
Формат стандартного имени:
[PAM]-{SERVER_HOSTNAME}-{HOSTNAME}

	or

[PAM]-{HOSTNAME}
*/
func getDefaultName() string {
	hostname, _ := os.Hostname()
	if serverHostname, ok := os.LookupEnv(hostEnvKey); ok {
		hostname = fmt.Sprintf("%s-%s", serverHostname, hostname)
	}
	hostRune := []rune("[PAM] - " + hostname)
	if len(hostRune) <= defaultNameMaxLen {
		return string(hostRune)
	}
	name := make([]rune, defaultNameMaxLen)
	index := defaultNameMaxLen / 2
	copy(name[:index], hostRune[:index])
	start := len(hostRune) - index
	copy(name[index:], hostRune[start:])
	return string(name)
}
