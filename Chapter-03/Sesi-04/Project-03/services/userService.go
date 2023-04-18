package services

import (
	"errors"
	"fmt"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/helpers"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/models"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/repository"
)

type UserService interface {
	Register(userReqData models.UserRegisterReq) (*models.UserRegisterRes, error)
	Login(userReqData models.UserLoginReq) (*string, error)
	GetProfile(userID string) (models.User, error)
}

type UserServiceIml struct {
	userRepository  repository.UserRepository
	photoRepository repository.PhotoRepository
	socalRepository repository.SocialRepository
}

func NewUserService(userRepo repository.UserRepository, photoRepo repository.PhotoRepository, socialRepo repository.SocialRepository) UserService {
	return &UserServiceIml{
		userRepository:  userRepo,
		photoRepository: photoRepo,
		socalRepository: socialRepo,
	}
}

func (s *UserServiceIml) Register(userReqData models.UserRegisterReq) (*models.UserRegisterRes, error) {
	userID := helpers.GenerateID()
	hashedPassword, err := helpers.Hash(userReqData.Password)
	if err != nil {
		return nil, err
	}

	newUser := models.User{
		UserID:   userID,
		Username: userReqData.Username,
		Email:    userReqData.Email,
		Password: hashedPassword,
		Age:      userReqData.Age,
	}

	err = s.userRepository.Create(newUser)
	if err != nil {
		return nil, err
	}

	return &models.UserRegisterRes{
		UserID:   newUser.UserID,
		Username: newUser.Username,
		Email:    newUser.Email,
		Password: newUser.Password,
		Age:      newUser.Age,
	}, nil
}

func (s *UserServiceIml) Login(userReqData models.UserLoginReq) (*string, error) {
	userResponse, err := s.userRepository.FindByUsername(userReqData.Username)
	if err != nil {
		return nil, err
	}

	isMatch := helpers.PasswordIsMatch(userReqData.Password, userResponse.Password)
	if isMatch == false {
		return nil, errors.New(fmt.Sprintf("Invalid username or password"))
	}

	token, err := helpers.GenerateToken(userResponse)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (s *UserServiceIml) GetProfile(userID string) (models.User, error) {
	user, err := s.userRepository.FindByID(userID)
	if err != nil {
		return models.User{}, err
	}

	photos, err := s.photoRepository.FindByUserID(userID)
	if err != nil {
		return models.User{}, err
	}

	socials, err := s.socalRepository.FindByUserID(userID)
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		UserID:      userID,
		Username:    user.Username,
		Email:       user.Email,
		Age:         user.Age,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		Photos:      photos,
		SocialMedia: socials,
	}, nil
}