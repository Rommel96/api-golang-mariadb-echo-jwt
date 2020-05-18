package auth

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte(key),
	}
	return middleware.JWTWithConfig(config)
}
