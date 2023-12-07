package config

import (
	"os"

	"github.com/HongXiangZuniga/login-go/pkg/http/rest/handler"
	"github.com/gin-gonic/gin"
)

var (
	ginEngine      *gin.Engine
	ginStatusGroup *gin.RouterGroup
)

func configHttp() {
	ginEngine = configGinEngine()
	ginStatusGroup = configStatusGroup()
	configStatusHandler()
}

func configGinEngine() *gin.Engine {
	env := os.Getenv("ENV")
	var engine *gin.Engine
	if env == "prod" {
		engine = gin.New()
		gin.SetMode(gin.ReleaseMode)
	} else {
		engine = gin.Default()
	}
	return engine
}

func GetEngine() *gin.Engine {
	return ginEngine
}

func configStatusGroup() *gin.RouterGroup {
	return ginEngine.Group("/status")
}

func configStatusHandler() {
	statusHandler := handler.NewStatusHandler()
	statusHandler.RegisterHandler(ginStatusGroup)
}
