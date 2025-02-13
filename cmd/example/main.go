package main

import (
	"context"
	"log"

	appLog "example/pkg/log"

	"example/internal/app"
	"example/internal/config"
	"example/internal/domain/example"
	"example/internal/domain/user"

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

var logOptions = appLog.DefaultOptions()

func init() {
	logger := appLog.NewLogger(logOptions).WithModule("metadata")
	logger.Infof("Version: %s", Version)
	logger.Infof("Commit: %s", Commit)
	logger.Infof("Build Time: %s", BuildTime)
}

func main() {
	server := fx.New(
		appLog.WithOptions(logOptions),
		fx.WithLogger(func(logger appLog.FxLogger) fxevent.Logger {
			return logger
		}),

		config.Module,
		app.Module,

		// Add your domain module here
		example.Module,
		user.Module,

		fx.Invoke(app.Start),
	)

	if err := server.Start(context.Background()); err != nil {
		log.Fatal(err)
	}

	<-server.Done()
}
