package middleware

import (
	"errors"
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

/*
	func (port *port) CheckUserSession() gin.HandlerFunc {
		return func(ctx *gin.Context) {
			if ctx.Request.URL.Path == LOGIN_URL {
				hash, _ := ctx.Cookie("session")
				authorize, err := port.authorizeService.Authorize(hash)
				if err != nil && err.Error() != "redis: nil" {
					port.logger.Error(err.Error())
				}
				if *authorize {
					ctx.Redirect(http.StatusMovedPermanently, LANDING_URL)
					return
				}
			} else {
				hash, _ := ctx.Cookie("session")
				authorize, err := port.authorizeService.Authorize(hash)
				if err != nil {
					port.logger.Error(err.Error())
				}
				if !*authorize {
					ctx.Redirect(http.StatusMovedPermanently, LOGIN_URL)
					return
				}
			}
			ctx.Next()
		}
	}
*/
func (port *port) CheckUserSession() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hash, err := ctx.Cookie("session")
		if err != nil && !errors.Is(err, http.ErrNoCookie) {
			port.logger.Error(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		isLoginURL := ctx.Request.URL.Path == LOGIN_URL
		authorize, err := port.authorizeService.Authorize(hash)
		if err != nil {
			port.logger.Error(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		if isLoginURL && *authorize {
			ctx.Redirect(http.StatusMovedPermanently, LANDING_URL)
			ctx.AbortWithStatus(http.StatusMovedPermanently)
			return
		}
		if !isLoginURL && !*authorize {
			ctx.Redirect(http.StatusMovedPermanently, LOGIN_URL)
			ctx.AbortWithStatus(http.StatusMovedPermanently)
			return
		}
		ctx.Next()
	}
}
