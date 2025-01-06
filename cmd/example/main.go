package main

import (
	"context"
	"example/internal/app"
	"example/internal/config"
	"example/internal/domain/user"
	"fmt"
	"log"

	"go.uber.org/fx"
)

// Build information. These will be populated by ldflags during build
var (
	Version   = "dev"
	Commit    = "none"
	BuildTime = "unknown"
)

func init() {
	fmt.Printf("Version:    %s\n", Version)
	fmt.Printf("Commit:     %s\n", Commit)
	fmt.Printf("Build Time: %s\n", BuildTime)
}

func main() {
	app := fx.New(
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
