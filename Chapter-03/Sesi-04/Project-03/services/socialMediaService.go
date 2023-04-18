package services

import (
	"errors"
	"time"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/helpers"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/models"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/repository"
)

type SocialService interface {
	Create(socialReqData models.SocialCreateReq, userID string) (*models.SocialResponse, error)
	GetAll() ([]models.SocialResponse, error)
	GetOne(socialID string) (models.SocialResponse, error)
	UpdateSocialMedia(socialReqData models.SocialUpdateReq, userID string, socialID string) (*models.SocialResponse, error)
	Delete(socialID string, userID string) (models.SocialResponse, error)
}

type SocialServiceIml struct {
	socialRepository repository.SocialRepository
}

func NewSocialService(socialRepo repository.SocialRepository) SocialService {
	return &SocialServiceIml{
		socialRepository: socialRepo,
	}
}

func (s *SocialServiceIml) Create(socialReqData models.SocialCreateReq, userID string) (*models.SocialResponse, error) {
	socialID := helpers.GenerateID()
	newSocial := models.SocialMedia{
		ID:             socialID,
		Name:           socialReqData.Name,
		SocialMediaURL: socialReqData.SocialMediaURL,
		UserID:         userID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err := s.socialRepository.Create(newSocial)
	if err != nil {
		return nil, err
	}

	return &models.SocialResponse{
		ID:             newSocial.ID,
		Name:           newSocial.Name,
		SocialMediaURL: newSocial.SocialMediaURL,
		UserID:         newSocial.UserID,
		CreatedAt:      newSocial.CreatedAt,
		UpdatedAt:      newSocial.UpdatedAt,
	}, nil
}

func (s *SocialServiceIml) GetAll() ([]models.SocialResponse, error) {
	photosResult, err := s.socialRepository.FindAll()
	if err != nil {
		return []models.SocialResponse{}, err
	}

	socialsResponse := []models.SocialResponse{}
	for _, socialRes := range photosResult {
		socialsResponse = append(socialsResponse, models.SocialResponse(socialRes))
	}

	return socialsResponse, nil
}

func (s *SocialServiceIml) GetOne(socialID string) (models.SocialResponse, error) {
	socialsResult, err := s.socialRepository.FindByID(socialID)
	if err != nil {
		return models.SocialResponse{}, err
	}

	return models.SocialResponse(socialsResult), nil
}

func (s *SocialServiceIml) UpdateSocialMedia(socialReqData models.SocialUpdateReq, userID string, socialID string) (*models.SocialResponse, error) {
	findSocialMediaResponse, err := s.socialRepository.FindByID(socialID)
	if err != nil {
		return nil, err
	}

	if userID != findSocialMediaResponse.UserID {
		return nil, errors.New("Unauthorized")
	}

	updatedSocialReq := models.SocialMedia{
		ID:             socialID,
		Name:           socialReqData.Name,
		SocialMediaURL: socialReqData.SocialMediaURL,
		UserID:         userID,
		UpdatedAt:      time.Now(),
	}

	err = s.socialRepository.Update(updatedSocialReq)
	if err != nil {
		return nil, err
	}

	return &models.SocialResponse{
		ID: socialID,
	}, nil
}

func (s *SocialServiceIml) Delete(socialID string, userID string) (models.SocialResponse, error) {
	findSocialResponse, err := s.socialRepository.FindByID(socialID)
	if err != nil {
		return models.SocialResponse{}, err
	}

	if userID != findSocialResponse.UserID {
		return models.SocialResponse{}, errors.New("Unauthorized")
	}

	err = s.socialRepository.Delete(models.SocialMedia{ID: socialID})
	if err != nil {
		return models.SocialResponse{}, err
	}

	return models.SocialResponse{
		ID: socialID,
	}, nil
}