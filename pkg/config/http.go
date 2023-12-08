package config

import (
	"os"

	"github.com/HongXiangZuniga/login-go/pkg/http/rest/handler"
	"github.com/gin-gonic/gin"
)

var (
	ginEngine      *gin.Engine
	ginStatusGroup *gin.RouterGroup
	ginViewGroup   *gin.RouterGroup
)

func configHttp() {
	ginEngine = configGinEngine()
	ginStatusGroup = configStatusGroup()
	ginViewGroup = configViewGroup()
	configStatusHandler()
	configViewHandler()
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
	engine.LoadHTMLGlob("pkg/template/*.html")
	engine.Static("/static", "./pkg/static")
	return engine
}

func GetEngine() *gin.Engine {
	return ginEngine
}

func configStatusGroup() *gin.RouterGroup {
	return ginEngine.Group("/status")
}

func configViewGroup() *gin.RouterGroup {
	return ginEngine.Group("/view")
}

func configStatusHandler() {
	statusHandler := handler.NewStatusHandler()
	statusHandler.RegisterHandler(ginStatusGroup)
}

func configViewHandler() {
	viewhander := handler.NewViewHanlder()
	viewhander.RegisterHandler(ginViewGroup)
}
