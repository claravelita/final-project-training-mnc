package usecase

import "github.com/claravelita/final-project-training-mnc/dtos"

type (
	UserUsecase interface {
		RegisterUser(request dtos.UserRequest) (dtos.JSONResponses, error)
		LoginUser(request dtos.LoginRequest) (dtos.JSONResponses, error)
		UpdateUser(userID int, request dtos.UserUpdateRequest) (dtos.JSONResponses, error)
		DeleteUser(userID int) (dtos.JSONResponses, error)
	}
)
