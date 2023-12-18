package config

import (
	"os"

	"github.com/HongXiangZuniga/login-go/pkg/http/rest/handler"
	"github.com/gin-gonic/gin"
)

var (
	ginEngine       *gin.Engine
	ginStatusGroup  *gin.RouterGroup
	ginViewGroup    *gin.RouterGroup
	ginAuthGroup    *gin.RouterGroup
	ginProfileGroup *gin.RouterGroup
)

func configHttp() {
	ginEngine = configGinEngine()
	ginStatusGroup = configStatusGroup()
	ginViewGroup = configViewGroup()
	ginAuthGroup = configAuthGroup()
	ginProfileGroup = configProfileGroup()
	configStatusHandler()
	configViewHandler()
	configAuthHandler()
	configProfileHandler()
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
	engine.Static("/static/css", "./pkg/static/css")
	engine.Static("/static/js", "./pkg/static/js")
	return engine
}

func GetEngine() *gin.Engine {
	return ginEngine
}

func configStatusGroup() *gin.RouterGroup {
	return ginEngine.Group("/status")
}

func configViewGroup() *gin.RouterGroup {
	return ginEngine.Group("/")
}

func configAuthGroup() *gin.RouterGroup {
	return ginEngine.Group("/auth")
}
func configProfileGroup() *gin.RouterGroup {
	return ginEngine.Group("/profile")
}

func configStatusHandler() {
	statusHandler := handler.NewStatusHandler()
	statusHandler.RegisterHandler(ginStatusGroup)
}

func configViewHandler() {
	viewhander := handler.NewViewHanlder()
	viewhander.RegisterHandler(ginViewGroup)
}

func configAuthHandler() {
	authHandler := handler.NewAuthHanler(autService)
	authHandler.RegisterHandler(ginAuthGroup)
}
func configProfileHandler() {
	ProfileHandler := handler.NewProfileHandler()
	ProfileHandler.RegisterHandler(ginProfileGroup)
}
