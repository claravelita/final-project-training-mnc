package repository

import (
	"errors"

	"github.com/claravelita/final-project-training-mnc/domain"
	"github.com/claravelita/final-project-training-mnc/dtos"
	"gorm.io/gorm"
)

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return socialMediaRepository{
		db,
	}
}

func (r socialMediaRepository) Create(socialMedia domain.SocialMedia) (domain.SocialMedia, error) {
	err := r.db.Create(&socialMedia).Error
	if err != nil {
		return domain.SocialMedia{}, err
	}

	return socialMedia, nil
}

func (r socialMediaRepository) GetAll(userID int) ([]domain.SocialMedia, error) {
	var all []domain.SocialMedia
	err := r.db.Preload("Users").Where("user_id = ?", userID).Find(&all).Error
	if err != nil {
		return []domain.SocialMedia{}, err
	}

	return all, nil
}

func (r socialMediaRepository) GetByID(socialMediaID int) (*domain.SocialMedia, error) {
	var socialMediaEntity domain.SocialMedia
	err := r.db.First(&socialMediaEntity, socialMediaID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &socialMediaEntity, err
}

func (r socialMediaRepository) Update(id int, request dtos.SocialMediaRequest) (domain.SocialMedia, error) {
	currentSocialMedia, err := r.GetByID(id)
	if err != nil {
		return domain.SocialMedia{}, err
	}

	updateSocialMedia := map[string]interface{}{
		"Name":           request.Name,
		"SocialMediaURL": request.SocialMediaURL,
	}

	err = r.db.Model(&currentSocialMedia).Updates(&updateSocialMedia).Error

	return *currentSocialMedia, err
}

func (r socialMediaRepository) Delete(id int) error {
	return r.db.Unscoped().Delete(&domain.SocialMedia{}, id).Error
}
