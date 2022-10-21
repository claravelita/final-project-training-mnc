package user

import (
	"strconv"

	"github.com/claravelita/final-project-training-mnc/common/command"
	"github.com/claravelita/final-project-training-mnc/domain"
	"github.com/claravelita/final-project-training-mnc/dtos"
)

func (u userImplementation) RegisterUser(request dtos.UserRequest) (dtos.JSONResponses, error) {
	password, err := u.auth.HashPassword(request.Password)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	checkUsername, err := u.repo.GetUserByUsername(request.Username)
	if checkUsername != nil {
		return command.BadRequestResponses("Username is Already Registered"), nil
	}
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	checkEmail, err := u.repo.GetUserByEmail(request.Email)
	if checkEmail != nil {
		return command.BadRequestResponses("Email is Already Registered"), nil
	}
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	requestUser := domain.User{
		Username: request.Username,
		Email:    request.Email,
		Password: password,
		Age:      request.Age,
	}
	newUser, err := u.repo.Create(requestUser)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	data := dtos.UserResponse{
		Username: newUser.Username,
		Email:    newUser.Email,
		Age:      newUser.Age,
		ID:       newUser.ID,
	}

	return command.SuccessCreatedResponses(data), nil
}

func (u userImplementation) LoginUser(request dtos.LoginRequest) (dtos.JSONResponses, error) {
	findEmail, err := u.repo.GetUserByEmail(request.Email)
	if findEmail == nil {
		return command.BadRequestResponses("Email not registered"), nil
	}
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	checkPassword := u.auth.CheckPasswordHash(request.Password, findEmail.Password)
	if checkPassword == true {
		token, tokenErr := u.auth.CreateJWTTokenLogin(strconv.FormatInt(int64(findEmail.ID), 10), findEmail.Email, findEmail.Username)
		if tokenErr != nil {
			return command.InternalServerResponses("Internal Server Error", err), err
		}
		data := dtos.TokenResponse{
			Token: "Bearer " + token,
		}
		return command.SuccessResponses(data), nil
	}

	return command.BadRequestResponses("Email or Password is incorrect"), nil
}

func (u userImplementation) UpdateUser(userID int, request dtos.UserUpdateRequest) (dtos.JSONResponses, error) {
	checkUser, err := u.repo.GetUserByID(userID)
	if checkUser == nil {
		return command.BadRequestResponses("User not found"), nil

	}
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	checkUsername, err := u.repo.GetUserByUsername(request.Username)
	if checkUsername != nil && checkUsername.ID != userID {
		return command.BadRequestResponses("Username is already taken"), nil
	}
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	checkEmail, err := u.repo.GetUserByEmail(request.Email)
	if checkEmail != nil && checkEmail.ID != userID {
		return command.BadRequestResponses("Email is already taken"), nil
	}
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	newUser, err := u.repo.Update(userID, request)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	data := dtos.UserResponse{
		Username: newUser.Username,
		Email:    newUser.Email,
		Age:      newUser.Age,
		ID:       newUser.ID,
		UpdateAt: newUser.AuditTable.UpdatedAt,
	}

	return command.SuccessCreatedResponses(data), nil
}

func (u userImplementation) DeleteUser(userID int) (dtos.JSONResponses, error) {
	checkUser, err := u.repo.GetUserByID(userID)
	if checkUser == nil {
		return command.BadRequestResponses("User not found"), nil
	}
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	err = u.repo.Delete(userID)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	message := dtos.MessageResponses{
		Message: "Your account has been successfully deleted",
	}

	return command.SuccessCreatedResponses(message), nil
}
