package handler

import "github.com/gin-gonic/gin"

type GinHandler interface {
	RegisterHandler(router *gin.RouterGroup)
}
