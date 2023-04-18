package repository

import (
	"errors"
	"fmt"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/models"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(commentReqData models.Comment) error
	FindAll() ([]models.Comment, error)
	FindByID(commentID string) (models.Comment, error)
	FindByPhotoID(photoID string) ([]models.Comment, error)
	Update(commentReqData models.Comment) error
	Delete(commentReqData models.Comment) error
}

type CommentRepositoryImpl struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &CommentRepositoryImpl{
		DB: db,
	}
}

func (r *CommentRepositoryImpl) Create(commentReqData models.Comment) error {
	err := r.DB.Create(&commentReqData).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *CommentRepositoryImpl) FindAll() ([]models.Comment, error) {
	comments := []models.Comment{}

	err := r.DB.Find(&comments).Error
	if err != nil {
		return []models.Comment{}, err
	}

	return comments, nil
}

func (r *CommentRepositoryImpl) FindByID(commentID string) (models.Comment, error) {
	comment := models.Comment{}

	err := r.DB.Debug().Where("comment_id = ?", commentID).Take(&comment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Comment{}, err
		}

		return models.Comment{}, err
	}

	return comment, nil
}

func (r *CommentRepositoryImpl) FindByPhotoID(photoID string) ([]models.Comment, error) {
	comments := []models.Comment{}

	err := r.DB.Where("photo_id = ?", photoID).Find(&comments).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []models.Comment{}, err
		}

		return []models.Comment{}, err
	}

	fmt.Println("comments: ", comments)

	return comments, nil
}

func (r *CommentRepositoryImpl) Update(commentReqData models.Comment) error {
	err := r.DB.Save(&models.Comment{
		CommentID: commentReqData.CommentID,
		Message:   commentReqData.Message,
		PhotoID:   commentReqData.PhotoID,
		UserID:    commentReqData.UserID,
		UpdatedAt: commentReqData.UpdatedAt,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *CommentRepositoryImpl) Delete(commentReqData models.Comment) error {
	err := r.DB.Delete(&commentReqData).Error
	if err != nil {
		return err
	}

	return nil
}