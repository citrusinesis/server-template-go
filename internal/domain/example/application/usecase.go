package application

import "example/internal/domain/example/persistence"

type ExampleUsecase struct {
	repo *persistence.ExampleRepository
}

func NewExampleUsecase(repo *persistence.ExampleRepository) *ExampleUsecase {
	return &ExampleUsecase{repo: repo}
}
