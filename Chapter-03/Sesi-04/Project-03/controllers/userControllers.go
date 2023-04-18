package controllers

import (
	"errors"
	"net/http"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/helpers"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/models"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetProfile(ctx *gin.Context)
}

type UserControllerImpl struct {
	userService services.UserService
}

func NewUserController(services services.UserService) UserController {
	return &UserControllerImpl{
		userService: services,
	}
}

// Register godoc
//
// @Summary		register user
// @Description	filled some form for registration
// @Tags		User
// @Accept		json
// @Produce		json
// @Param		request	body		models.UserRegisterReq	true	"request is required"
// @Success		200		{object}	models.SuccessResponse{data=models.UserRegisterRes}
// @Failure		400		{object}	models.ErrorResponse{errors=interface{}}
// @Failure		500		{object}	models.ErrorResponse{errors=interface{}}
// @Router		/register [post]
func (c *UserControllerImpl) Register(ctx *gin.Context) {
	request := models.UserRegisterReq{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})
		return
	}

	validateErrs := []error{}
	validateErrs = helpers.UserRegisterValidator(request)
	if validateErrs != nil {
		errResponseStr := make([]string, len(validateErrs))
		for i, err := range validateErrs {
			errResponseStr[i] = err.Error()
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: errResponseStr,
		})
		return
	}

	response, err := c.userService.Register(request)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
				Code:   http.StatusBadRequest,
				Status: "Bad Request",
				Errors: errors.New("This email or username already registered").Error(),
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.SuccessResponse{
		Code:    http.StatusOK,
		Message: "User registered successfully",
		Data:    response,
	})
}

// Login godoc
//
// @Summary		login user
// @Description	login user using username and password
// @Tags		User
// @Accept		json
// @Produce		json
// @Param		request	body		models.UserLoginReq	true	"request is required"
// @Success		200		{object}	models.SuccessResponse{data=models.UserLoginRes}
// @Failure		400		{object}	models.ErrorResponse{errors=interface{}}
// @Failure		500		{object}	models.ErrorResponse{errors=interface{}}
// @Router		/login [post]
func (c *UserControllerImpl) Login(ctx *gin.Context) {
	request := models.UserLoginReq{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})
		return
	}

	validateErrs := []error{}
	validateErrs = helpers.UserLoginValidator(request)
	if validateErrs != nil {
		errResponseStr := make([]string, len(validateErrs))
		for i, err := range validateErrs {
			errResponseStr[i] = err.Error()
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: errResponseStr,
		})
		return
	}

	response, err := c.userService.Login(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.SuccessResponse{
		Code:    http.StatusOK,
		Message: "User login successfully",
		Data: models.UserLoginRes{
			Token: *response,
		},
	})
}

func (c *UserControllerImpl) GetProfile(ctx *gin.Context) {
	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: "UserID doesn't exist",
		})
		return
	}

	user, err := c.userService.GetProfile(userID.(string))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.UserResponse{
		UserID:      user.UserID,
		Username:    user.Username,
		Email:       user.Email,
		Age:         user.Age,
		Photos:      user.Photos,
		SocialMedia: user.SocialMedia,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	})
}