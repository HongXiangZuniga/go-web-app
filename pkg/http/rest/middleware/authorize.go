package middleware

import (
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
		if ctx.Request.URL.Path == LOGIN_URL {
			hash, _ := ctx.Cookie("session")
			authorize, err := port.authorizeService.Authorize(hash)
			if err != nil {
				port.logger.Error(err.Error())
			}
			if *authorize {
				ctx.Redirect(301, LANDING_URL)
			}
			ctx.Next()
			return
		} else {
			hash, _ := ctx.Cookie("session")
			authorize, err := port.authorizeService.Authorize(hash)
			if err != nil {
				port.logger.Error(err.Error())
			}
			if !*authorize {
				ctx.Redirect(301, LOGIN_URL)
			}
		}
		ctx.Next()
		return
	}
}
