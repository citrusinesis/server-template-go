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
	// Add middleware here

	e.GET("/", HealthCheck)

	return e
}

func NewServer(config *config.Config, e *echo.Echo) *http.Server {

	return &http.Server{
		Addr:    config.Server.String(),
		Handler: e,
	}
}

func Start(lifecycle fx.Lifecycle, server *http.Server, logger appLog.Logger, config *config.Config) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.ListenAndServe()

			logger.WithModule("app").
				WithFields(appLog.Fields{"addr": config.Server.String()}).
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
)
