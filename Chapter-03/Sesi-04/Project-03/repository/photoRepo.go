package repository

import (
	"errors"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/models"
	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(photoReqData models.Photo) error
	FindAll() ([]models.Photo, error)
	FindByID(photoID string) (models.Photo, error)
	FindByUserID(userID string) ([]models.Photo, error)
	Update(photoReqData models.Photo) error
	Delete(photoReqData models.Photo) error
}

type PhotoRepositoryImpl struct {
	DB *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &PhotoRepositoryImpl{
		DB: db,
	}
}

func (r *PhotoRepositoryImpl) Create(photoReqData models.Photo) error {
	err := r.DB.Create(&photoReqData).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *PhotoRepositoryImpl) FindAll() ([]models.Photo, error) {
	photos := []models.Photo{}

	err := r.DB.Find(&photos).Error
	if err != nil {
		return []models.Photo{}, err
	}

	return photos, nil
}

func (r *PhotoRepositoryImpl) FindByID(photoID string) (models.Photo, error) {
	photo := models.Photo{}

	err := r.DB.Debug().Where("photo_id = ?", photoID).Take(&photo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Photo{}, err
		}

		return models.Photo{}, err
	}

	return photo, nil
}

func (r *PhotoRepositoryImpl) FindByUserID(userID string) ([]models.Photo, error) {
	photos := []models.Photo{}

	err := r.DB.Where("user_id = ?", userID).Find(&photos).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []models.Photo{}, err
		}

		return []models.Photo{}, err
	}

	return photos, nil
}

func (r *PhotoRepositoryImpl) Update(photoReqData models.Photo) error {
	err := r.DB.Save(&models.Photo{
		PhotoID:   photoReqData.PhotoID,
		Title:     photoReqData.Title,
		Caption:   photoReqData.Caption,
		PhotoURL:  photoReqData.PhotoURL,
		UserID:    photoReqData.UserID,
		UpdatedAt: photoReqData.UpdatedAt,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *PhotoRepositoryImpl) Delete(photoReqData models.Photo) error {
	err := r.DB.Delete(&photoReqData).Error
	if err != nil {
		return err
	}

	return nil
}