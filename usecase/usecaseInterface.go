package usecase

import "github.com/claravelita/final-project-training-mnc/dtos"

type (
	UserUsecase interface {
		RegisterUser(request dtos.UserRequest) (dtos.JSONResponses, error)
		LoginUser(request dtos.LoginRequest) (dtos.JSONResponses, error)
		UpdateUser(userID int, request dtos.UserUpdateRequest) (dtos.JSONResponses, error)
		DeleteUser(userID int) (dtos.JSONResponses, error)
	}

	PhotoUsecase interface {
		CreatePhoto(userID int, request dtos.PhotoRequest) (dtos.JSONResponses, error)
		GetAllPhoto(userID int) (dtos.JSONResponses, error)
		UpdatePhoto(userID, photoID int, request dtos.PhotoRequest) (dtos.JSONResponses, error)
		DeletePhoto(userID, photoID int) (dtos.JSONResponses, error)
	}

	CommentUsecase interface {
	}
)
