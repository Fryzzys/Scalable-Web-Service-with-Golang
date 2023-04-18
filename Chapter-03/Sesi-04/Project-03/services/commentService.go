package services

import (
	"errors"
	"time"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/helpers"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/models"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/repository"
)

type CommentService interface {
	Create(commentReqData models.CommentCreateReq, userID string, photoID string) (*models.CommentResponse, error)
	GetAll() ([]models.CommentResponse, error)
	GetOne(commentID string) (models.CommentResponse, error)
	UpdateComment(commentReqData models.CommentUpdateReq, userID string, commentID string) (*models.CommentResponse, error)
	Delete(commentID string, userID string) (models.CommentDeleteRes, error)
}

type CommentServiceIml struct {
	commentRepository repository.CommentRepository
	photoRepository   repository.PhotoRepository
}

func NewCommentService(commentRepo repository.CommentRepository, photoRepo repository.PhotoRepository) CommentService {
	return &CommentServiceIml{
		commentRepository: commentRepo,
		photoRepository:   photoRepo,
	}
}

func (s *CommentServiceIml) Create(commentReqData models.CommentCreateReq, userID string, photoID string) (*models.CommentResponse, error) {
	_, err := s.photoRepository.FindByID(photoID)
	if err != nil {
		return nil, err
	}

	commentID := helpers.GenerateID()
	newComment := models.Comment{
		CommentID: commentID,
		Message:   commentReqData.Message,
		PhotoID:   photoID,
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.commentRepository.Create(newComment)
	if err != nil {
		return nil, err
	}

	return &models.CommentResponse{
		CommentID: newComment.CommentID,
		Message:   newComment.Message,
		PhotoID:   newComment.PhotoID,
		UserID:    newComment.UserID,
		CreatedAt: newComment.CreatedAt,
		UpdatedAt: newComment.UpdatedAt,
	}, nil
}

func (s *CommentServiceIml) GetAll() ([]models.CommentResponse, error) {
	commentsResult, err := s.commentRepository.FindAll()
	if err != nil {
		return []models.CommentResponse{}, err
	}

	commentsResponse := []models.CommentResponse{}
	for _, commentRes := range commentsResult {
		commentsResponse = append(commentsResponse, models.CommentResponse(commentRes))
	}

	return commentsResponse, nil
}

func (s *CommentServiceIml) GetOne(commentID string) (models.CommentResponse, error) {
	commentResult, err := s.commentRepository.FindByID(commentID)
	if err != nil {
		return models.CommentResponse{}, err
	}

	return models.CommentResponse(commentResult), nil
}

func (s *CommentServiceIml) UpdateComment(commentReqData models.CommentUpdateReq, userID string, commentID string) (*models.CommentResponse, error) {
	findCommentResponse, err := s.commentRepository.FindByID(commentID)
	if err != nil {
		return nil, err
	}

	if userID != findCommentResponse.UserID {
		return nil, errors.New("Unauthorized")
	}

	updatedCommentReq := models.Comment{
		CommentID: findCommentResponse.CommentID,
		Message:   commentReqData.Message,
		PhotoID:   findCommentResponse.PhotoID,
		UserID:    userID,
		UpdatedAt: time.Now(),
	}

	err = s.commentRepository.Update(updatedCommentReq)
	if err != nil {
		return nil, err
	}

	return &models.CommentResponse{
		CommentID: commentID,
	}, nil
}

func (s *CommentServiceIml) Delete(commentlID string, userID string) (models.CommentDeleteRes, error) {
	findCommentResponse, err := s.commentRepository.FindByID(commentlID)
	if err != nil {
		return models.CommentDeleteRes{}, err
	}

	if userID != findCommentResponse.UserID {
		return models.CommentDeleteRes{}, errors.New("Unauthorized")
	}

	err = s.commentRepository.Delete(models.Comment{CommentID: commentlID})
	if err != nil {
		return models.CommentDeleteRes{}, err
	}

	return models.CommentDeleteRes{
		CommentID: commentlID,
	}, nil
}