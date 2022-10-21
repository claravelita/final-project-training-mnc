package socialmedia

import (
	"time"

	"github.com/claravelita/final-project-training-mnc/common/command"
	"github.com/claravelita/final-project-training-mnc/domain"
	"github.com/claravelita/final-project-training-mnc/dtos"
)

func (c socialMediaImplementation) CreateSocialMedia(userID int, request dtos.SocialMediaRequest) (dtos.JSONResponses, error) {
	checkUser, err := c.repoUser.GetUserByID(userID)
	if checkUser == nil {
		return command.BadRequestResponses("User not found"), nil
	}
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	requestSocialMedia := domain.SocialMedia{
		Name:           request.Name,
		SocialMediaURL: request.SocialMediaURL,
		UserID:         userID,
	}
	newSocialMedia, err := c.repo.Create(requestSocialMedia)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	timeNow := time.Now()
	data := dtos.SocialMediaResponse{
		ID:             newSocialMedia.ID,
		Name:           newSocialMedia.Name,
		SocialMediaURL: newSocialMedia.SocialMediaURL,
		UserID:         newSocialMedia.UserID,
		CreatedAt:      &timeNow,
	}

	return command.SuccessCreatedResponses(data), nil
}

func (c socialMediaImplementation) GetAllSocialMedia(userID int) (dtos.JSONResponses, error) {
	checkUser, err := c.repoUser.GetUserByID(userID)
	if checkUser == nil {
		return command.BadRequestResponses("User not found"), nil
	}
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	getSocialMedia, err := c.repo.GetAll(userID)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	timeNow := time.Now()
	str := "string"
	var socialMedias []dtos.SocialMediaResponse
	for _, sm := range getSocialMedia {
		socialMedia := dtos.SocialMediaResponse{
			ID: sm.ID,
			Name: sm.Name,
			SocialMediaURL: sm.SocialMediaURL,
			UserID: sm.UserID,
			CreatedAt:&timeNow,
			Users: &dtos.UserResponse{
				ID: sm.Users.ID,
				Username: sm.Users.Username,
				Email: sm.Users.Email,
				Age: sm.Users.Age,
				UpdateAt: &sm.Users.AuditTable.UpdatedAt,
				ProfileImageURL: &str,
			},
		}
		socialMedias = append(socialMedias,socialMedia)
	}
	
	return command.SuccessResponses(socialMedias), nil
}

func (u socialMediaImplementation) UpdateSocialMedia(userID, socialMediaID int, request dtos.SocialMediaRequest) (dtos.JSONResponses, error) {
	checkUser, err := u.repoUser.GetUserByID(userID)
	if checkUser == nil {
		return command.BadRequestResponses("User not found"), nil
	}
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	checkSocialMedia, err := u.repo.GetByID(socialMediaID)
	if checkSocialMedia == nil {
		return command.BadRequestResponses("User not found"), nil
	}
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	updateSocialMedia, err := u.repo.Update(socialMediaID, request)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	timeNow := time.Now()
	data := dtos.SocialMediaResponse{
		ID:             updateSocialMedia.ID,
		Name:           updateSocialMedia.Name,
		SocialMediaURL: updateSocialMedia.SocialMediaURL,
		UserID:         updateSocialMedia.UserID,
		UpdatedAt:      &timeNow,
	}

	return command.SuccessResponses(data), nil
}

func (u socialMediaImplementation) DeleteSocialMedia(userID, socialMediaID int) (dtos.JSONResponses, error) {
	checkUser, err := u.repoUser.GetUserByID(userID)
	if checkUser == nil {
		return command.BadRequestResponses("User not found"), nil
	}
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	checkSocialMedia, err := u.repo.GetByID(socialMediaID)
	if checkSocialMedia == nil {
		return command.BadRequestResponses("Social Media not found"), nil
	}
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	if checkUser.ID != checkSocialMedia.UserID{
		return command.BadRequestResponses("Social Media user is incorrect"), nil
	}

	err = u.repo.Delete(socialMediaID)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	message := dtos.MessageResponses{
		Message: "Your social media has been successfully deleted",
	}

	return command.SuccessResponses(message), nil
}













