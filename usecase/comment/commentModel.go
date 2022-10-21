package comment

import (
	"net/http"

	"github.com/claravelita/final-project-training-mnc/common/command"
	"github.com/claravelita/final-project-training-mnc/domain"
	"github.com/claravelita/final-project-training-mnc/dtos"
)

func (c commentImplementation) CreateComment(userID int, request dtos.CommentRequest) (dtos.JSONResponses, error) {
	checkUser, err := c.repoUser.GetUserByID(userID)
	if checkUser == nil {
		return command.BadRequestResponses("User not found"), nil
	}
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	checkPhoto, err := c.repoPhoto.GetPhotoByID(request.PhotoID)
	if checkPhoto == nil {
		return command.BadRequestResponses("Photo not found"), nil
	}
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	requestComment := domain.Comment{
		UserID:  userID,
		PhotoID: request.PhotoID,
		Message: request.Message,
	}
	newComment, err := c.repo.Create(requestComment)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	data := dtos.CommentResponse{
		ID:        newComment.ID,
		PhotoID:   newComment.PhotoID,
		Message:   newComment.Message,
		UserID:    newComment.UserID,
		CreatedAt: &newComment.AuditTable.CreatedAt,
	}

	return command.SuccessCreatedResponses(data), nil
}

func (c commentImplementation) GetAllComment(userID int) (dtos.JSONResponses, error) {
	checkUser, err := c.repoUser.GetUserByID(userID)
	if checkUser == nil {
		return command.BadRequestResponses("User not found"), nil
	}
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	getComment, err := c.repo.GetAll(userID)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	var comments []dtos.CommentResponse
	for _, c := range getComment {
		comment := dtos.CommentResponse{
			ID:        c.ID,
			UserID:    c.UserID,
			PhotoID:   c.PhotoID,
			Message:   c.Message,
			CreatedAt: &c.AuditTable.CreatedAt,
			UpdatedAt: &c.AuditTable.UpdatedAt,
			Users: &dtos.UserResponse{
				ID:       c.Users.ID,
				Username: c.Users.Username,
				Email:    c.Users.Email,
			},
			Photos: &dtos.PhotoResponse{
				ID:       c.Photos.ID,
				Title:    c.Photos.Title,
				Caption:  c.Photos.Caption,
				PhotoURL: c.Photos.PhotoURL,
				UserID:   c.Photos.UserID,
			},
		}

		comments = append(comments, comment)

	}

	return command.SuccessResponses(comments), nil
}

func (u commentImplementation) UpdateComment(userID, commentID int, request dtos.CommentUpdateRequest) (dtos.JSONResponses, error) {
	checkIDs := u.IDCommentValidation(userID, commentID)
	if checkIDs.Code != http.StatusOK {
		return checkIDs, nil
	}

	updateComment, err := u.repo.Update(commentID, request)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	data := dtos.CommentResponse{
		ID:        updateComment.ID,
		UserID:    updateComment.UserID,
		PhotoID:   updateComment.PhotoID,
		Message:   updateComment.Message,
		UpdatedAt: &updateComment.AuditTable.UpdatedAt,
	}

	return command.SuccessCreatedResponses(data), nil
}

func (u commentImplementation) DeleteComment(userID, commentID int) (dtos.JSONResponses, error) {
	checkIDs := u.IDCommentValidation(userID, commentID)
	if checkIDs.Code != http.StatusOK {
		return checkIDs, nil
	}

	err := u.repo.Delete(commentID)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err), err
	}

	message := dtos.MessageResponses{
		Message: "Your comment has been successfully deleted",
	}

	return command.SuccessResponses(message), nil
}

func (u commentImplementation) IDCommentValidation(userID, commentID int) dtos.JSONResponses {
	checkUser, err := u.repoUser.GetUserByID(userID)
	if checkUser == nil {
		return command.BadRequestResponses("User not found")

	}
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err)
	}

	checkComment, err := u.repo.GetCommentByID(commentID)
	if checkComment == nil {
		return command.BadRequestResponses("Photo not found")
	}
	if err != nil {
		return command.InternalServerResponses("Internal Server Error", err)
	}

	if checkComment.UserID != userID {
		return command.BadRequestResponses("Comment user is incorrect")
	}

	return command.SuccessResponses(nil)
}
