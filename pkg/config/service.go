package config

import (
	"os"
	"strings"

	"github.com/HongXiangZuniga/login-go/pkg/auth"
	"github.com/HongXiangZuniga/login-go/pkg/authorize"
)

var (
	autService       auth.Service
	authorizeService authorize.Service
)

func configService() {
	authorizeService = configAuthorize()
	autService = configAuth()
}

func configAuth() auth.Service {
	isDemo := os.Getenv("DEMO")
	if strings.ToLower(isDemo) == "true" {
		return auth.NewService(AuthSQLRepo, logger, true, os.Getenv("DUMMY_USER"), os.Getenv("DUMMY_PASSWORD"))
	} else {
		return auth.NewService(AuthSQLRepo, logger, false, "", "")
	}
}

func configAuthorize() authorize.Service {
	return authorize.NewService(AuthorizeRedisRepository, logger)
}
