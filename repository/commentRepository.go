package repository

import "gorm.io/gorm"

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return commentRepository{
		db,
	}
}
