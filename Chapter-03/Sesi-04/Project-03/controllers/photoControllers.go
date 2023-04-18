package controllers

import (
	"net/http"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/helpers"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/models"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/services"
	"github.com/gin-gonic/gin"
)

type PhotoController interface {
	CreatePhoto(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	UpdatePhoto(ctx *gin.Context)
	DeletePhoto(ctx *gin.Context)
}

type PhotoControllerImpl struct {
	photoService services.PhotoService
}

func NewPhotoController(services services.PhotoService) PhotoController {
	return &PhotoControllerImpl{
		photoService: services,
	}
}

// CreatePhoto godoc
//
//	@Summary		create photo
//	@Description	create photo for a particular user
//	@Tags			Photo
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.PhotoCreateReq	true	"request is required"
//	@Success		200		{object}	models.SuccessResponse{data=models.PhotoCreateRes}
//	@Failure		400		{object}	models.ErrorResponse{errors=interface{}}
//	@Failure		500		{object}	models.ErrorResponse{errors=interface{}}
//	@Security		BearerAuth
//	@Router			/photos [post]
func (c *PhotoControllerImpl) CreatePhoto(ctx *gin.Context) {
	var request models.PhotoCreateReq

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: "UserID doesn't exist",
		})
		return
	}

	validateErrs := []error{}
	validateErrs = helpers.PhotoCreateValidator(request)
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

	response, err := c.photoService.Create(request, userID.(string))
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
		Message: "Photo created successfully",
		Data:    response,
	})
}

// GetAllPhoto godoc
//
//	@Summary		get all photo
//	@Description	get all photo
//	@Tags			Photo
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	models.SuccessResponse{data=[]models.PhotoResponse}
//	@Failure		400		{object}	models.ErrorResponse{errors=interface{}}
//	@Failure		500		{object}	models.ErrorResponse{errors=interface{}}
//	@Security		BearerAuth
//	@Router			/photos [get]
func (c *PhotoControllerImpl) GetAll(ctx *gin.Context) {
	response, err := c.photoService.GetAll()
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
		Message: "Get all photo successfully",
		Data:    response,
	})
}

// GetOnePhoto godoc
//
//	@Summary		get one photo
//	@Description	get one photo
//	@Tags			Photo
//	@Accept			json
//	@Produce		json
//  @Param          photo_id   path      string  true  "PhotoID"
//	@Success		200		{object}	models.SuccessResponse{data=models.PhotoResponse}
//	@Failure		400		{object}	models.ErrorResponse{errors=interface{}}
//	@Failure		500		{object}	models.ErrorResponse{errors=interface{}}
//	@Security		BearerAuth
//	@Router			/photos/{photo_id} [get]
func (c *PhotoControllerImpl) GetOne(ctx *gin.Context) {
	photoID := ctx.Param("photo_id")

	response, err := c.photoService.GetOne(photoID)
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
		Message: "Get photo successfully",
		Data:    response,
	})
}

// UpdatePhoto godoc
//
//	@Summary		update photo
//	@Description	update photo
//	@Tags			Photo
//	@Accept			json
//	@Produce		json
//  @Param          photo_id   path      string  true  "PhotoID"
//	@Param			request	body		models.PhotoUpdateReq	true	"request is required"
//	@Success		200		{object}	models.SuccessResponse{data=models.PhotoUpdateRes}
//	@Failure		400		{object}	models.ErrorResponse{errors=interface{}}
//	@Failure		500		{object}	models.ErrorResponse{errors=interface{}}
//	@Security		BearerAuth
//	@Router			/photos/{photo_id} [put]
func (c *PhotoControllerImpl) UpdatePhoto(ctx *gin.Context) {
	var request models.PhotoUpdateReq
	photoID := ctx.Param("photo_id")

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: "UserID doesn't exist",
		})
		return
	}

	validateErrs := []error{}
	validateErrs = helpers.PhotoUpdateValidator(request)
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

	response, err := c.photoService.UpdatePhoto(request, userID.(string), photoID)
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
		Message: "Photo updated successfully",
		Data: models.PhotoUpdateRes{
			PhotoID: response.PhotoID,
		},
	})
}

// DeletePhoto godoc
//
//	@Summary		delete photo
//	@Description	delete photo
//	@Tags			Photo
//	@Accept			json
//	@Produce		json
//  @Param          photo_id   path      string  true  "PhotoID"
//	@Success		200		{object}	models.SuccessResponse{data=models.PhotoDeleteRes}
//	@Failure		400		{object}	models.ErrorResponse{errors=interface{}}
//	@Failure		500		{object}	models.ErrorResponse{errors=interface{}}
//	@Security		BearerAuth
//	@Router			/photos/{photo_id} [delete]
func (c *PhotoControllerImpl) DeletePhoto(ctx *gin.Context) {
	photoID := ctx.Param("photo_id")

	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: "UserID doesn't exist",
		})
		return
	}

	response, err := c.photoService.Delete(photoID, userID.(string))
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
		Message: "Photo deleted successfully",
		Data: models.PhotoUpdateRes{
			PhotoID: response.PhotoID,
		},
	})
}