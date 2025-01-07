package application

import "example/internal/domain/user/persistence"

type UserExampleUsecase struct {
	repo *persistence.UserRepository
}

func NewUserExampleUsecase(repo *persistence.UserRepository) *UserExampleUsecase {
	return &UserExampleUsecase{repo: repo}
}
