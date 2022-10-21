package repository

import (
	"errors"

	"github.com/claravelita/final-project-training-mnc/domain"
	"github.com/claravelita/final-project-training-mnc/dtos"
	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return commentRepository{
		db,
	}
}

func (r commentRepository) Create(comment domain.Comment) (domain.Comment, error) {
	err := r.db.Create(&comment).Error
	if err != nil {
		return domain.Comment{}, err
	}

	return comment, nil
}

func (r commentRepository) GetAll(userID int) ([]domain.Comment, error) {
	var all []domain.Comment
	err := r.db.Preload("Users").Preload("Photos").Where("user_id = ?", userID).Find(&all).Error
	if err != nil {
		return []domain.Comment{}, err
	}

	return all, nil
}

func (r commentRepository) GetCommentByID(commentID int) (*domain.Comment, error) {
	var commentEntity domain.Comment
	err := r.db.First(&commentEntity, commentID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &commentEntity, err
}

func (r commentRepository) Update(id int, request dtos.CommentUpdateRequest) (domain.Comment, error) {
	currentComment, err := r.GetCommentByID(id)
	if err != nil {
		return domain.Comment{}, err
	}

	updateComment := map[string]interface{}{
		"Message": request.Message,
	}

	err = r.db.Model(&currentComment).Updates(&updateComment).Error

	return *currentComment, err
}

func (r commentRepository) Delete(id int) error {
	return r.db.Unscoped().Delete(&domain.Comment{}, id).Error
}
