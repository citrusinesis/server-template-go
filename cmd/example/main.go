package main

import (
	"context"
	"log"

	"example/internal/app"
	"example/internal/config"
	"example/internal/domain/user"

	"go.uber.org/fx"
)

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
