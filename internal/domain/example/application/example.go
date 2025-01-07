package application

import (
	"example/internal/core/usecase"
	"example/internal/domain/example/persistence"
)

type ExampleInput struct{}

func (t ExampleInput) Validate() error {
	return nil
}

var _ usecase.Usecase[ExampleInput, string] = (*ExampleUsecase)(nil)

type ExampleUsecase struct {
	repo *persistence.ExampleRepository
}

func NewExampleUsecase(repo *persistence.ExampleRepository) *ExampleUsecase {
	return &ExampleUsecase{repo: repo}
}

func (u *ExampleUsecase) Execute(input ExampleInput) (string, error) {
	return "example", nil
}
