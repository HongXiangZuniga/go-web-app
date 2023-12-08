package handler

import "github.com/gin-gonic/gin"

type Viewhandler struct {
}

func NewViewHanlder() GinHandler {
	return &Viewhandler{}
}

func (impl *Viewhandler) RegisterHandler(router *gin.RouterGroup) {
	router.GET("/login", impl.Get)
}

func (impl *Viewhandler) Get(ctx *gin.Context) {
	ctx.JSON(204, nil)
}
