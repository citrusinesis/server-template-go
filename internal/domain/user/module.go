package user

import (
	"example/internal/domain/user/application"
	"example/internal/domain/user/persistence"
	"example/internal/domain/user/presentation"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		// Add your repository implementation here
		persistence.NewUserRepository,

		// Add your usecase here
		application.NewUserExampleUsecase,
	),
	fx.Invoke(presentation.NewUserHandler),
)
