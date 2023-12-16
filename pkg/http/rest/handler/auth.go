package handler

import (
	"github.com/HongXiangZuniga/login-go/pkg/auth"
	http "github.com/HongXiangZuniga/login-go/pkg/http/rest"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService auth.Service
}

func NewAuthHanler(authService auth.Service) GinHandler {
	return &AuthHandler{authService: authService}
}

func (impl *AuthHandler) RegisterHandler(router *gin.RouterGroup) {
	router.POST("", impl.GetAuth)
}

func (impl *AuthHandler) GetAuth(ctx *gin.Context) {
	var loginData http.LoginRequest
	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, err := impl.authService.Authorization(loginData.User, loginData.Password)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if *result {
		ctx.JSON(200, result)
	} else {
		ctx.JSON(401, gin.H{"error": "User Unauthorized"})
	}

}
