package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
}

func NewProfileHandler() GinHandler {
	return &ProfileHandler{}
}

func (impl *ProfileHandler) RegisterHandler(router *gin.RouterGroup) {
	router.GET("", impl.GetProfile)
}

func (impl *ProfileHandler) GetProfile(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "profile.html", gin.H{})
}
