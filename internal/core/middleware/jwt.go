package middleware

import (
	"example/internal/config"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWT(cfg *config.Config, skipPaths ...string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		Skipper: func(c echo.Context) bool {
			for _, s := range skipPaths {
				if c.Path() == s {
					return true
				}
			}
			return false
		},
		SigningKey: []byte(cfg.Server.JWTSecret),
	})
}
