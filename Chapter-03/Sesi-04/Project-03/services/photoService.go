package services

import (
	"errors"
	"time"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/helpers"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/models"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/repository"
)

type PhotoService interface {
	Create(photoReqData models.PhotoCreateReq, userID string) (*models.PhotoCreateRes, error)
	GetAll() ([]models.PhotoResponse, error)
	GetOne(photoID string) (models.PhotoResponse, error)
	UpdatePhoto(photoReqData models.PhotoUpdateReq, userID string, photoID string) (*models.PhotoResponse, error)
	Delete(photoID string, userID string) (models.PhotoResponse, error)
}

type PhotoServiceIml struct {
	photoRepository   repository.PhotoRepository
	commentRepository repository.CommentRepository
}

func NewPhotoService(photoRepo repository.PhotoRepository, commentRepo repository.CommentRepository) PhotoService {
	return &PhotoServiceIml{
		photoRepository:   photoRepo,
		commentRepository: commentRepo,
	}
}

func (s *PhotoServiceIml) Create(photoReqData models.PhotoCreateReq, userID string) (*models.PhotoCreateRes, error) {
	photoID := helpers.GenerateID()
	newPhoto := models.Photo{
		PhotoID:   photoID,
		Title:     photoReqData.Title,
		Caption:   photoReqData.Caption,
		PhotoURL:  photoReqData.PhotoURL,
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.photoRepository.Create(newPhoto)
	if err != nil {
		return nil, err
	}

	return &models.PhotoCreateRes{
		PhotoID:   newPhoto.PhotoID,
		Title:     newPhoto.Title,
		Caption:   newPhoto.Caption,
		PhotoURL:  newPhoto.PhotoURL,
		UserID:    newPhoto.UserID,
		CreatedAt: newPhoto.CreatedAt,
		UpdatedAt: newPhoto.UpdatedAt,
	}, nil
}

func (s *PhotoServiceIml) GetAll() ([]models.PhotoResponse, error) {
	photosResult, err := s.photoRepository.FindAll()
	if err != nil {
		return []models.PhotoResponse{}, err
	}

	photosResponse := []models.PhotoResponse{}
	for _, photoRes := range photosResult {
		photosResponse = append(photosResponse, models.PhotoResponse(photoRes))
	}

	return photosResponse, nil
}

func (s *PhotoServiceIml) GetOne(photoID string) (models.PhotoResponse, error) {
	photosResult, err := s.photoRepository.FindByID(photoID)
	if err != nil {
		return models.PhotoResponse{}, err
	}

	comments := []models.Comment{}
	commentsResponse, err := s.commentRepository.FindByPhotoID(photoID)
	for _, comment := range commentsResponse {
		comments = append(comments, models.Comment(comment))
	}
	if err != nil {
		return models.PhotoResponse{}, err
	}

	return models.PhotoResponse{
		PhotoID:   photosResult.PhotoID,
		Title:     photosResult.Title,
		Caption:   photosResult.Caption,
		PhotoURL:  photosResult.PhotoURL,
		UserID:    photosResult.UserID,
		Comments:  comments,
		CreatedAt: photosResult.CreatedAt,
		UpdatedAt: photosResult.UpdatedAt,
	}, nil
}

func (s *PhotoServiceIml) UpdatePhoto(photoReqData models.PhotoUpdateReq, userID string, photoID string) (*models.PhotoResponse, error) {
	findPhotoResponse, err := s.photoRepository.FindByID(photoID)
	if err != nil {
		return nil, err
	}

	if userID != findPhotoResponse.UserID {
		return nil, errors.New("Unauthorized")
	}

	updatedPhotoReq := models.Photo{
		PhotoID:   photoID,
		Title:     photoReqData.Title,
		Caption:   photoReqData.Caption,
		PhotoURL:  photoReqData.PhotoURL,
		UserID:    userID,
		UpdatedAt: time.Now(),
	}

	err = s.photoRepository.Update(updatedPhotoReq)
	if err != nil {
		return nil, err
	}

	return &models.PhotoResponse{
		PhotoID: photoID,
	}, nil
}

func (s *PhotoServiceIml) Delete(photoID string, userID string) (models.PhotoResponse, error) {
	findPhotoResponse, err := s.photoRepository.FindByID(photoID)
	if err != nil {
		return models.PhotoResponse{}, err
	}

	if userID != findPhotoResponse.UserID {
		return models.PhotoResponse{}, errors.New("Unauthorized")
	}

	err = s.photoRepository.Delete(models.Photo{PhotoID: photoID})
	if err != nil {
		return models.PhotoResponse{}, err
	}

	return models.PhotoResponse{
		PhotoID: photoID,
	}, nil
}