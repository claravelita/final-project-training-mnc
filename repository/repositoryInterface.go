package repository

import (
	"github.com/claravelita/final-project-training-mnc/domain"
	"github.com/claravelita/final-project-training-mnc/dtos"
)

type (
	UserRepository interface {
		GetUserByUsername(username string) (*domain.User, error)
		GetUserByEmail(email string) (*domain.User, error)
		GetUserByID(userID int) (*domain.User, error)
		Create(user domain.User) (domain.User, error)
		Update(id int, request dtos.UserUpdateRequest) (domain.User, error)
		Delete(id int) error
	}

	PhotoRepository interface {
		Create(photo domain.Photo) (domain.Photo, error)
		GetAll(userID int) ([]domain.Photo, error)
		GetPhotoByID(userID int) (*domain.Photo, error)
		Update(id int, request dtos.PhotoRequest) (domain.Photo, error)
		Delete(id int) error
	}
)
