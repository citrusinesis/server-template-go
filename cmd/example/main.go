package main

import (
	"context"
	"log"

	"example/internal/app"
	"example/internal/config"
	"example/internal/domain/user"
	appLog "example/pkg/log"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

// Build information. These will be populated by ldflags during build
var (
	Name      = "example" // Change this to your app name
	Version   = "dev"
	Commit    = "none"
	BuildTime = "unknown"
)

func printVersion(logger appLog.Logger) {
	logger = logger.WithModule("startup")
	logger.Infof("Version: %s", Version)
	logger.Infof("Commit: %s", Commit)
	logger.Infof("Build Time: %s", BuildTime)
}

func main() {
	app := fx.New(
		appLog.WithOptions(&appLog.Options{
			FormatterType: appLog.TextFormatter,
			FilePath:      "",
			Level:         appLog.DebugLevel,
		}),
		fx.WithLogger(func(logger appLog.FxLogger) fxevent.Logger {
			return logger
		}),

		config.Module,
		app.Module,

		// Add your domain module here
		user.Module,

		fx.Invoke(printVersion, app.Start),
	)

	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}

	<-app.Done()
}
