package config

import (
	"fmt"
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
		fmt.Println("DEMO ACTIVATE")
		fmt.Println("USER:" + os.Getenv("DUMMY_USER"))
		fmt.Println("PASSWORD:" + os.Getenv("DUMMY_PASSWORD"))
		return auth.NewService(true, os.Getenv("DUMMY_USER"), os.Getenv("DUMMY_PASSWORD"))
	} else {
		return auth.NewService(false, "", "")
	}
}
