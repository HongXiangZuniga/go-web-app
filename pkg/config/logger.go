package config

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	err    error
)

func configLogger() {
	level := os.Getenv("LOG_LEVEL")
	config := zap.NewProductionConfig()
	switch strings.ToLower(level) {
	case "debug":
		config.Level.SetLevel(zap.DebugLevel)
	case "info":
		config.Level.SetLevel(zap.InfoLevel)
	case "warning":
		config.Level.SetLevel(zap.WarnLevel)
	case "error":
		config.Level.SetLevel(zap.ErrorLevel)
	default:
		level = "error"
		config.Level.SetLevel(zap.ErrorLevel)
	}
	fmt.Println("[CONFIG] LOG_LEVEL CONFIGURED: " + level)
	logger, err = config.Build()
	if err != nil {
		panic(err)
	}
}

func GetLogger() *zap.Logger {
	return logger
}
