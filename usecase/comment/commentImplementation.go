package comment

import (
	"github.com/claravelita/final-project-training-mnc/common/helper"
	"github.com/claravelita/final-project-training-mnc/repository"
	"github.com/claravelita/final-project-training-mnc/usecase"
)

type commentImplementation struct {
	repo      repository.CommentRepository
	repoPhoto repository.PhotoRepository
	repoUser  repository.UserRepository
	auth      helper.AuthHelper
}

func NewCommentImplementation(repo repository.CommentRepository, repoPhoto repository.PhotoRepository, repoUser repository.UserRepository, auth helper.AuthHelper) usecase.CommentUsecase {
	return &commentImplementation{
		repo:      repo,
		repoPhoto: repoPhoto,
		repoUser:  repoUser,
		auth:      auth,
	}
}
