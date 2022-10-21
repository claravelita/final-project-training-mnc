package user

import (
	"github.com/claravelita/final-project-training-mnc/common/helper"
	"github.com/claravelita/final-project-training-mnc/repository"
	"github.com/claravelita/final-project-training-mnc/usecase"
)

type userImplementation struct {
	repo repository.UserRepository
	auth helper.AuthHelper
}

func NewUserImplementation(repo repository.UserRepository, auth helper.AuthHelper) usecase.UserUsecase {
	return &userImplementation{
		repo: repo,
		auth: auth,
	}
}
