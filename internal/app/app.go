package app

import (
	"context"
	"net/http"

	"example/internal/config"
	appMiddleware "example/internal/core/middleware"
	appLog "example/pkg/log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
)

func NewEcho(cfg *config.Config, logger appLog.Logger) *echo.Echo {
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.Recover())
	e.Use(appMiddleware.Logger(logger))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper: func(c echo.Context) bool {
			return false
		},
		AllowOrigins:     cfg.Server.CORS.AllowOrigins,
		AllowMethods:     cfg.Server.CORS.AllowMethods,
		AllowHeaders:     cfg.Server.CORS.AllowHeaders,
		ExposeHeaders:    cfg.Server.CORS.ExposeHeaders,
		AllowCredentials: cfg.Server.CORS.AllowCredentials,
		MaxAge:           cfg.Server.CORS.MaxAge,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:X-XSRF-TOKEN",
	}))

	e.Use(middleware.BodyLimitWithConfig(middleware.BodyLimitConfig{
		Skipper: func(c echo.Context) bool {
			return false
		},
		Limit: cfg.Server.MaxBodySize,
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	// Add middleware here

	return e
}

func NewServer(config *config.Config, e *echo.Echo) *http.Server {
	return &http.Server{
		Addr:    config.Server.Bind.String(),
		Handler: e,
	}
}

func Start(lifecycle fx.Lifecycle, server *http.Server, logger appLog.Logger, config *config.Config) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.ListenAndServe()

			logger.WithModule("app").
				WithFields(appLog.Fields{"addr": config.Server.Bind.String()}).
				Info("Listening")

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}

var Module = fx.Module("app",
	fx.Provide(
		NewEcho,
		NewServer,
	),
	fx.Invoke(
		NewHealthCheckHandler,
	),
)
