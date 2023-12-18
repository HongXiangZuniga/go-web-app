package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Viewhandler struct {
}

func NewViewHanlder() GinHandler {
	return &Viewhandler{}
}

func (impl *Viewhandler) RegisterHandler(router *gin.RouterGroup) {
	router.GET("/login", impl.GetLogin)
}

func (impl *Viewhandler) GetLogin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}

