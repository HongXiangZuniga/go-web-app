package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func Config() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not load .env")
	}
	configLogger()
	configDataBase()
	configRepository()
	configService()
	configMiddleware()
	configHttp()
}
