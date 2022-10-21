package repository

import (
	"errors"

	"github.com/claravelita/final-project-training-mnc/domain"
	"github.com/claravelita/final-project-training-mnc/dtos"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		db,
	}
}

func (r userRepository) GetUserByUsername(username string) (*domain.User, error) {
	var userEntity domain.User
	err := r.db.Where("username = ?", username).First(&userEntity).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &userEntity, err
}

func (r userRepository) GetUserByEmail(email string) (*domain.User, error) {
	var userEntity domain.User
	err := r.db.Where("email = ?", email).First(&userEntity).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &userEntity, err
}

func (r userRepository) GetUserByID(userID int) (*domain.User, error) {
	var userEntity domain.User
	err := r.db.First(&userEntity, userID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &userEntity, err
}

func (r userRepository) Create(user domain.User) (domain.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r userRepository) Update(id int, request dtos.UserUpdateRequest) (domain.User, error) {
	currentUser, err := r.GetUserByID(id)
	if err != nil {
		return domain.User{}, err
	}

	updateUser := map[string]interface{}{
		"Email":    request.Email,
		"Username": request.Username,
	}

	err = r.db.Model(&currentUser).Updates(&updateUser).Error

	return *currentUser, err
}

func (r userRepository) Delete(id int) error {
	return r.db.Unscoped().Delete(&domain.User{}, id).Error
}
