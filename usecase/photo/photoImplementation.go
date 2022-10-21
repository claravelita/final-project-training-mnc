package photo

import (
	"github.com/claravelita/final-project-training-mnc/common/helper"
	"github.com/claravelita/final-project-training-mnc/repository"
	"github.com/claravelita/final-project-training-mnc/usecase"
)

type photoImplementation struct {
	repo     repository.PhotoRepository
	repoUser repository.UserRepository
	auth     helper.AuthHelper
}

func NewPhotoImplementation(repo repository.PhotoRepository, repoUser repository.UserRepository, auth helper.AuthHelper) usecase.PhotoUsecase {
	return &photoImplementation{
		repo:     repo,
		repoUser: repoUser,
		auth:     auth,
	}
}
