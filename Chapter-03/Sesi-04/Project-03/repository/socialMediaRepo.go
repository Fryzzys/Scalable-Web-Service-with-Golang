package repository

import (
	"errors"
	"project-myGram/models"
	"gorm.io/gorm"
)

type SocialRepository interface {
	Create(photoReqData models.SocialMedia) error
	FindAll() ([]models.SocialMedia, error)
	FindByID(socialID string) (models.SocialMedia, error)
	FindByUserID(userID string) ([]models.SocialMedia, error)
	Update(socialReqData models.SocialMedia) error
	Delete(photoReqData models.SocialMedia) error
}

type SocialRepositoryImpl struct {
	DB *gorm.DB
}

func NewSocialRepository(db *gorm.DB) SocialRepository {
	return &SocialRepositoryImpl{
		DB: db,
	}
}

func (r *SocialRepositoryImpl) Create(socialReqData models.SocialMedia) error {
	err := r.DB.Create(&socialReqData).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *SocialRepositoryImpl) FindAll() ([]models.SocialMedia, error) {
	socials := []models.SocialMedia{}

	err := r.DB.Find(&socials).Error
	if err != nil {
		return []models.SocialMedia{}, err
	}

	return socials, nil
}

func (r *SocialRepositoryImpl) FindByID(socialID string) (models.SocialMedia, error) {
	social := models.SocialMedia{}

	err := r.DB.Debug().Where("id = ?", socialID).Take(&social).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.SocialMedia{}, err
		}

		return models.SocialMedia{}, err
	}

	return social, nil
}

func (r *SocialRepositoryImpl) FindByUserID(userID string) ([]models.SocialMedia, error) {
	socials := []models.SocialMedia{}

	err := r.DB.Debug().Where("user_id = ?", userID).Find(&socials).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []models.SocialMedia{}, err
		}

		return []models.SocialMedia{}, err
	}

	return socials, nil
}

func (r *SocialRepositoryImpl) Update(socialReqData models.SocialMedia) error {
	err := r.DB.Save(&models.SocialMedia{
		ID:             socialReqData.ID,
		Name:           socialReqData.Name,
		SocialMediaURL: socialReqData.SocialMediaURL,
		UserID:         socialReqData.UserID,
		UpdatedAt:      socialReqData.UpdatedAt,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *SocialRepositoryImpl) Delete(socialReqData models.SocialMedia) error {
	err := r.DB.Delete(&socialReqData).Error
	if err != nil {
		return err
	}

	return nil
}