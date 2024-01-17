package middleware

import (
	"net/http"

	"github.com/HongXiangZuniga/login-go/pkg/authorize"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthorizeMiddlewareServices interface {
	CheckUserSession() gin.HandlerFunc
}

type port struct {
	authorizeService authorize.Service
	logger           *zap.Logger
}

func NewAuthorizeMiddlewareServices(authorize authorize.Service, logger *zap.Logger) AuthorizeMiddlewareServices {
	return &port{authorize, logger}
}

const (
	LOGIN_URL   = "/login"
	LANDING_URL = "/profile"
)

func (port *port) CheckUserSession() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		switch ctx.Request.URL.Path {
		case LOGIN_URL:
			hash, _ := ctx.Cookie("session")
			authorize, err := port.authorizeService.Authorize(hash)
			if err != nil && err.Error() != "redis: nil" {
				port.logger.Error(err.Error())
			}
			if hash == "" || (err != nil && err.Error() == "redis: nil") {
				ctx.Next()
				return
			}
			if *authorize {
				ctx.Redirect(http.StatusFound, LANDING_URL)
				ctx.Abort()
				return
			}
			ctx.Next()
		case "/auth", "/static" + ctx.Request.URL.Path:
			ctx.Next()
		default:
			hash, _ := ctx.Cookie("session")
			authorize, err := port.authorizeService.Authorize(hash)
			if err != nil {
				port.logger.Error(err.Error())
			}
			if !*authorize {
				ctx.Redirect(http.StatusFound, LOGIN_URL)
				ctx.Abort()
				return
			}
			ctx.Next()
		}
	}
}
