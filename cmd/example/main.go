package main

import (
	"context"
	"fmt"
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

func init() {
	fmt.Println("--------------------")
	fmt.Printf("Starting %s...\n", Name)
	fmt.Printf("Version:    %s\n", Version)
	fmt.Printf("Commit:     %s\n", Commit)
	fmt.Printf("Build Time: %s\n", BuildTime)
	fmt.Println("--------------------")
}

func main() {
	app := fx.New(
		appLog.Module,
		fx.WithLogger(func(logger appLog.FxLogger) fxevent.Logger {
			return logger
		}),

		config.Module,
		app.Module,

		// Add your domain module here
		user.Module,

		fx.Invoke(app.Start),
	)

	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}

	<-app.Done()
}
