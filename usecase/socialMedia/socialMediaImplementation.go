package socialmedia

import (
	"github.com/claravelita/final-project-training-mnc/common/helper"
	"github.com/claravelita/final-project-training-mnc/repository"
	"github.com/claravelita/final-project-training-mnc/usecase"
)

type socialMediaImplementation struct {
	repo        repository.SocialMediaRepository
	repoComment repository.CommentRepository
	repoPhoto   repository.PhotoRepository
	repoUser    repository.UserRepository
	auth        helper.AuthHelper
}

func NewSocialMediaImplementation(repo repository.SocialMediaRepository, repoComment repository.CommentRepository, repoPhoto repository.PhotoRepository, repoUser repository.UserRepository, auth helper.AuthHelper) usecase.SocialMediaUsecase {
	return &socialMediaImplementation{
		repo:        repo,
		repoComment: repoComment,
		repoPhoto:   repoPhoto,
		repoUser:    repoUser,
		auth:        auth,
	}
}
