package config

import (
	"os"
	"strings"

	"github.com/HongXiangZuniga/login-go/pkg/authentication"
	"github.com/HongXiangZuniga/login-go/pkg/authorize"
)

var (
	autService       authentication.Service
	authorizeService authorize.Service
)

func configService() {
	authorizeService = configAuthorize()
	autService = configAuth()
}

func configAuth() authentication.Service {
	isDemo := os.Getenv("DEMO")
	if strings.ToLower(isDemo) == "true" {
		return authentication.NewService(AuthSQLRepo, logger, true, os.Getenv("DUMMY_USER"), os.Getenv("DUMMY_PASSWORD"))
	} else {
		return authentication.NewService(AuthSQLRepo, logger, false, "", "")
	}
}

func configAuthorize() authorize.Service {
	return authorize.NewService(AuthorizeRedisRepository, logger)
}
