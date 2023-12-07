package handler

import (
	"github.com/gin-gonic/gin"
)

type statusHandler struct {
}

func NewStatusHandler() GinHandler {
	return &statusHandler{}
}

func (impl *statusHandler) RegisterHandler(router *gin.RouterGroup) {
	router.GET("/", impl.Get)
}

func (impl *statusHandler) Get(ctx *gin.Context) {
	ctx.JSON(204, nil)
}
