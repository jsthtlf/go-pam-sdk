package config

import (
	"fmt"
	"github.com/jsthtlf/go-pam-sdk/common"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

const (
	hostEnvKey = "SERVER_HOSTNAME"

	defaultNameMaxLen = 128
)

var config *Config

type Config struct {
	Name    string `mapstructure:"TERMINAL_NAME"`
	Comment string `mapstructure:"TERMINAL_COMMENT"`

	CoreHost       string `mapstructure:"CORE_HOST"`
	BootstrapToken string `mapstructure:"BOOTSTRAP_TOKEN"`
	LogLevel       string `mapstructure:"LOG_LEVEL"`
	LogFormat      string `mapstructure:"LOG_FORMAT"`
	LanguageCode   string `mapstructure:"LANGUAGE_CODE"`

	TerminalType string

	RootPath          string
	DataFolderPath    string
	LogDirPath        string
	KeyFolderPath     string
	AccessKeyFilePath string
	ReplayFolderPath  string
}

func GetCurrentConfig() Config {
	if config == nil {
		return getDefaultConfig()
	}
	return *config
}

func SetupConfig(configPath string) *Config {
	var conf = getDefaultConfig()
	loadConfigFromEnv(&conf)
	loadConfigFromFile(configPath, &conf)
	config = &conf
	return &conf
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
		if err := common.EnsureDirExist(folders[i]); err != nil {
			log.Fatalf("Create folder failed: %s", err)
		}
	}
	return Config{
		Name:    defaultName,
		Comment: "Pam terminal",

		CoreHost:       "localhost",
		BootstrapToken: "",
		LogLevel:       "INFO",
		LogFormat:      "%time% [%lvl%] %msg%",
		LanguageCode:   "ru",

		TerminalType: "PAM",

		AccessKeyFilePath: accessKeyFilePath,
		RootPath:          rootPath,
		DataFolderPath:    dataFolderPath,
		LogDirPath:        LogDirPath,
		KeyFolderPath:     keyFolderPath,
		ReplayFolderPath:  replayFolderPath,
	}
}

func have(path string) bool {
	_, err := os.Stat(path)
	return err == nil
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

func loadConfigFromFile(path string, conf *Config) {
	var err error
	if have(path) {
		fileViper := viper.New()
		fileViper.SetConfigFile(path)
		if err = fileViper.ReadInConfig(); err == nil {
			if err = fileViper.Unmarshal(conf); err == nil {
				log.Printf("Load config from %s: success\n", path)
				return
			}
		}
	}
	if err != nil {
		log.Fatalf("Load config from %s failed: %s\n", path, err)
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
	hostRune := []rune("[" + config.TerminalType + "] - " + hostname)
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
