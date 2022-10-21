package photo

import (
	"github.com/claravelita/final-project-training-mnc/common/command"
	"github.com/claravelita/final-project-training-mnc/domain"
	"github.com/claravelita/final-project-training-mnc/dtos"
)

func (p photoImplementation) CreatePhoto(userID int, request dtos.PhotoRequest) (dtos.JSONResponses, error) {
	requestUser := domain.Photo{
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoURL: request.PhotoURL,
		UserID:   userID,
	}
	newPhoto, err := p.repo.Create(requestUser)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	data := dtos.PhotoResponse{
		ID:        newPhoto.ID,
		Title:     newPhoto.Title,
		Caption:   newPhoto.Caption,
		PhotoURL:  newPhoto.PhotoURL,
		UserID:    newPhoto.UserID,
		CreatedAt: &newPhoto.AuditTable.CreatedAt,
	}

	return command.SuccessCreatedResponses(data), nil
}

func (p photoImplementation) GetAllPhoto(userID int) (dtos.JSONResponses, error) {
	getPhoto, err := p.repo.GetAll(userID)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	var photos []dtos.PhotoResponse
	for _, p := range getPhoto {
		photo := dtos.PhotoResponse{
			ID:        p.ID,
			Title:     p.Title,
			Caption:   p.Caption,
			PhotoURL:  p.PhotoURL,
			UserID:    p.UserID,
			CreatedAt: &p.AuditTable.CreatedAt,
			Users: &dtos.UserResponse{
				Email:    p.Users.Email,
				Username: p.Users.Username,
			},
		}

		photos = append(photos, photo)

	}

	return command.SuccessResponses(photos), nil
}

func (u photoImplementation) UpdatePhoto(userID, photoID int, request dtos.PhotoRequest) (dtos.JSONResponses, error) {
	checkUser, err := u.repoUser.GetUserByID(userID)
	if checkUser == nil {
		return command.BadRequestResponses("User not found"), nil

	}
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	checkPhoto, err := u.repo.GetPhotoByID(photoID)
	if checkPhoto == nil {
		return command.BadRequestResponses("Photo not found"), nil
	}
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	if checkPhoto.UserID != userID {
		return command.BadRequestResponses("Photo user is incorrect"), nil
	}

	newPhoto, err := u.repo.Update(photoID, request)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	data := dtos.PhotoResponse{
		ID:        newPhoto.ID,
		Title:     newPhoto.Title,
		Caption:   newPhoto.Caption,
		PhotoURL:  newPhoto.PhotoURL,
		UserID:    newPhoto.UserID,
		UpdatedAt: &newPhoto.AuditTable.UpdatedAt,
	}

	return command.SuccessCreatedResponses(data), nil
}

func (u photoImplementation) DeletePhoto(userID, photoID int) (dtos.JSONResponses, error) {
	checkUser, err := u.repoUser.GetUserByID(userID)
	if checkUser == nil {
		return command.BadRequestResponses("User not found"), nil
	}
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	checkPhoto, err := u.repo.GetPhotoByID(photoID)
	if checkPhoto == nil {
		return command.BadRequestResponses("Photo not found"), nil
	}
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	if checkPhoto.UserID != userID {
		return command.BadRequestResponses("Photo user is incorrect"), nil
	}

	err = u.repo.Delete(photoID)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	message := dtos.MessageResponses{
		Message: "Your photo has been successfully deleted",
	}

	return command.SuccessCreatedResponses(message), nil
}
