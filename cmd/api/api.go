package main

import (
	"github.com/HongXiangZuniga/login-go/pkg/config"
)

func main() {
	config.Config()
	gin := config.GetEngine()
	gin.Run()
	/*jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(jsonHandler)

	logger.Info("This is an Info message", slog.Int("version", 1.0)) // <
	*/
}
