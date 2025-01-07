package user

import (
	"example/internal/domain/example/application"
	"example/internal/domain/example/persistence"
	"example/internal/domain/example/presentation"

	"go.uber.org/fx"
)

var Module = fx.Module("example",
	fx.Provide(
		// Add your repository implementation here
		persistence.NewExampleRepository,

		// Add your usecase here
		application.NewExampleUsecase,
	),
	fx.Invoke(presentation.NewExampleHandler),
)
