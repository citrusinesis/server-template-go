package app

import (
	"context"

	"example/internal/config"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func NewEcho(cfg *config.Config) *echo.Echo {
	e := echo.New()
	// Add middleware here
	return e
}

func Start(lifecycle fx.Lifecycle, e *echo.Echo, cfg *config.Config) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go e.Start(cfg.Server.Port)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
}

var Module = fx.Options(
	fx.Provide(NewEcho),
)
