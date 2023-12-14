package config

import (
	"os"
	"strings"

	"github.com/HongXiangZuniga/login-go/pkg/auth"
)

var (
	autService auth.Service
)

func configService() {
	autService = configAuth()
}

func configAuth() auth.Service {
	isDemo := os.Getenv("DEMO")
	if strings.ToLower(isDemo) == "true" {
		return auth.NewService(logger, true, os.Getenv("DUMMY_USER"), os.Getenv("DUMMY_PASSWORD"))
	} else {
		return auth.NewService(logger, false, "", "")
	}
}
