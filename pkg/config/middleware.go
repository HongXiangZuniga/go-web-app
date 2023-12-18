package config

import "github.com/HongXiangZuniga/login-go/pkg/http/rest/middleware"

var (
	AuthorizeMiddlewareServices middleware.AuthorizeMiddlewareServices
)

func configMiddleware() {
	AuthorizeMiddlewareServices = configAuthorizeMiddlewareServices()
}

func configAuthorizeMiddlewareServices() middleware.AuthorizeMiddlewareServices {
	return middleware.NewAuthorizeMiddlewareServices(authorizeService, logger)
}
