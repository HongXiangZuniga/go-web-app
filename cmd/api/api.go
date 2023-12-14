package main

import (
	"github.com/HongXiangZuniga/login-go/pkg/config"
)

func main() {
	config.Config()
	gin := config.GetEngine()
	gin.Run()
}
