package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jsthtlf/go-pam-sdk/config"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()
var logLevels = map[string]logrus.Level{
	"DEBUG":    logrus.DebugLevel,
	"INFO":     logrus.InfoLevel,
	"WARN":     logrus.WarnLevel,
	"ERROR":    logrus.ErrorLevel,
	"FATAL":    logrus.FatalLevel,
	"CRITICAL": logrus.PanicLevel,
}

func Initial(conf *config.Config) {
	formatter := &Formatter{
		LogFormat:       "%time% [%lvl%] %msg%",
		TimestampFormat: "2006-01-02 15:04:05",
	}
	level, ok := logLevels[strings.ToUpper(conf.LogLevel)]
	if !ok {
		level = logrus.InfoLevel
	}

	if strings.Contains(conf.LogFormat, "%time%") &&
		strings.Contains(conf.LogFormat, "%lvl%") &&
		strings.Contains(conf.LogFormat, "%msg%") {
		formatter.LogFormat = conf.LogFormat
	}

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logger.SetFormatter(formatter)
	logger.SetOutput(os.Stdout)
	logger.SetLevel(level)

	// Output to file
	logFilePath := filepath.Join(conf.LogDirPath, "PAM.log")
	rotateFileHook, err := NewRotateFileHook(RotateFileConfig{
		Filename:   logFilePath,
		MaxSize:    50,
		MaxBackups: 7,
		MaxAge:     7,
		LocalTime:  true,
		Level:      level,
		Formatter:  formatter,
	})
	if err != nil {
		fmt.Printf("Create log rotate hook error: %s\n", err)
		return
	}
	logger.AddHook(rotateFileHook)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Panic(args ...interface{}) {
	logrus.Panic(args...)
}

func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}
