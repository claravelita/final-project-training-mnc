package repository

import (
	"errors"

	"github.com/claravelita/final-project-training-mnc/domain"
	"github.com/claravelita/final-project-training-mnc/dtos"
	"gorm.io/gorm"
)

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return photoRepository{
		db,
	}
}

func (r photoRepository) Create(photo domain.Photo) (domain.Photo, error) {
	err := r.db.Create(&photo).Error
	if err != nil {
		return domain.Photo{}, err
	}

	return photo, nil
}

func (r photoRepository) GetAll(userID int) ([]domain.Photo, error) {
	var all []domain.Photo
	err := r.db.Preload("Users").Where("user_id = ?", userID).Find(&all).Error
	if err != nil {
		return []domain.Photo{}, err
	}

	return all, nil
}

func (r photoRepository) GetPhotoByID(photoID int) (*domain.Photo, error) {
	var photoEntity domain.Photo
	err := r.db.First(&photoEntity, photoID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &photoEntity, err
}

func (r photoRepository) Update(id int, request dtos.PhotoRequest) (domain.Photo, error) {
	currentPhoto, err := r.GetPhotoByID(id)
	if err != nil {
		return domain.Photo{}, err
	}

	updatePhoto := map[string]interface{}{
		"Title":    request.Title,
		"Caption":  request.Caption,
		"PhotoURL": request.PhotoURL,
	}

	err = r.db.Model(&currentPhoto).Updates(&updatePhoto).Error

	return *currentPhoto, err
}

func (r photoRepository) Delete(id int) error {
	return r.db.Unscoped().Delete(&domain.Photo{}, id).Error
}
