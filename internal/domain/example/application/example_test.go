package application_test

import (
	"testing"

	"example/internal/core/usecase"
	"example/internal/domain/example/application"
	"example/internal/domain/example/persistence"
)

func TestExampleUsecase(t *testing.T) {
	repo := &persistence.ExampleRepository{}
	uc := application.NewExampleUsecase(repo)

	validator := func(expected, actual string) (bool, error) {
		return expected == actual, nil
	}
	tester := usecase.NewTester(uc, t, validator)
	tester.AddTestcase(
		usecase.Testcase[application.ExampleInput, string]{
			Input:    application.ExampleInput{},
			Expected: "example",
		},
	)

	tester.Run()
}
