package handler

import (
	"github.com/HongXiangZuniga/login-go/pkg/auth"
	"github.com/HongXiangZuniga/login-go/pkg/authorize"
	http "github.com/HongXiangZuniga/login-go/pkg/http/rest"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthHandler struct {
	authorizeService authorize.Service
	authService      auth.Service
	logger           *zap.Logger
}

func NewAuthHanler(AuthorizeService authorize.Service, authService auth.Service, logger *zap.Logger) GinHandler {
	return &AuthHandler{
		authorizeService: AuthorizeService,
		authService:      authService,
		logger:           logger}
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
		hash, err := impl.authorizeService.SetHash(loginData.User)
		if err != nil {
			impl.logger.Error(err.Error())
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.SetCookie("session", *hash, 300, "/", "localhost:8080", false, true)
		ctx.Redirect(301, "/profile")
	} else {
		ctx.JSON(401, gin.H{"error": "User Unauthorized"})
	}

}
