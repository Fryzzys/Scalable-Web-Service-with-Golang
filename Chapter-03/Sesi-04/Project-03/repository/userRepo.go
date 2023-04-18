package repository

import (
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(userReqData models.User) error
	FindByID(userID string) (models.User, error)
	FindByUsername(username string) (models.User, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (r *UserRepositoryImpl) Create(userReqData models.User) error {
	err := r.DB.Create(&userReqData).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryImpl) FindByID(userID string) (models.User, error) {
	var user models.User
	err := r.DB.First(&user, "user_id = ?", userID).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) FindByUsername(username string) (models.User, error) {
	var user models.User
	err := r.DB.First(&user, "username = ?", username).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}