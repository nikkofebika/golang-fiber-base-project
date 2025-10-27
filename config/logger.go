package config

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func InitLogger(cfg *AppConfig) {
	// logrus.New()

	// use logrus default formatter by default
	if strings.ToLower(cfg.LogFormat) == "json" {
		fmt.Println("SILITTTTT")
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02T15:04:05Z07:00",
		})
	}

	// use info level by default
	appEnv := strings.ToLower(cfg.AppEnv)
	switch appEnv {
	case "production":
		logrus.SetLevel(logrus.ErrorLevel)
	case "development":
		logrus.SetLevel(logrus.WarnLevel)
	}

	// Output: stdout + optional file
	var outputs []io.Writer
	outputs = append(outputs, os.Stdout)
	if cfg.LogFilePath != "" {
		f, err := os.OpenFile(cfg.LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
		if err == nil {
			outputs = append(outputs, f)
		} else {
			logrus.Warnf("cannot open log file %s: %v", cfg.LogFilePath, err)
		}
	}
	logrus.SetOutput(io.MultiWriter(outputs...))

	// return logger
}
