package controllers

import (
	"errors"
	"net/http"
	"project-myGram/helpers"
	"project-myGram/models"
	"project-myGram/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentController interface {
	CreateComment(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	UpdateComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)
}

type CommentControllerImpl struct {
	commentService services.CommentService
}

func NewCommentController(services services.CommentService) CommentController {
	return &CommentControllerImpl{
		commentService: services,
	}
}

// CreateComment godoc
//
//	@Summary		create comment
//	@Description	create comment for a particular user
//	@Tags			Comment
//	@Accept			json
//	@Produce		json
//  @Param          photo_id   path      string  true  "photo_id"
//	@Param			request	body		models.CommentCreateReq	true	"request is required"
//	@Success		200		{object}	models.SuccessResponse{data=models.CommentResponse}
//	@Failure		400		{object}	models.ErrorResponse{errors=interface{}}
//	@Failure		500		{object}	models.ErrorResponse{errors=interface{}}
//	@Security		BearerAuth
//	@Router			/comments/{photo_id} [post]
func (c *CommentControllerImpl) CreateComment(ctx *gin.Context) {
	var request models.CommentCreateReq
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
	validateErrs = helpers.CommentCreateValidator(request)
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

	response, err := c.commentService.Create(request, userID.(string), photoID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{
				Code:   http.StatusBadRequest,
				Status: "Bad Request",
				Errors: errors.New("Photo not found").Error(),
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
		Message: "Comment created successfully",
		Data:    response,
	})
}

// GetAllComment godoc
//
//	@Summary		get all comment
//	@Description	get all comment
//	@Tags			Comment
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	models.SuccessResponse{data=[]models.CommentResponse}
//	@Failure		400		{object}	models.ErrorResponse{errors=interface{}}
//	@Failure		500		{object}	models.ErrorResponse{errors=interface{}}
//	@Security		BearerAuth
//	@Router			/comments [get]
func (c *CommentControllerImpl) GetAll(ctx *gin.Context) {
	response, err := c.commentService.GetAll()
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
		Message: "Get all comment successfully",
		Data:    response,
	})
}

// GetOneComment godoc
//
//	@Summary		get one comment
//	@Description	get one comment
//	@Tags			Comment
//	@Accept			json
//	@Produce		json
//  @Param          comment_id   path      string  true  "comment_id"
//	@Success		200		{object}	models.SuccessResponse{data=models.CommentResponse}
//	@Failure		400		{object}	models.ErrorResponse{errors=interface{}}
//	@Failure		500		{object}	models.ErrorResponse{errors=interface{}}
//	@Security		BearerAuth
//	@Router			/comments/{comment_id} [get]
func (c *CommentControllerImpl) GetOne(ctx *gin.Context) {
	commentID := ctx.Param("comment_id")

	response, err := c.commentService.GetOne(commentID)
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
		Message: "Get comment successfully",
		Data:    response,
	})
}

// UpdateComment godoc
//
//	@Summary		update comment
//	@Description	update comment
//	@Tags			Comment
//	@Accept			json
//	@Produce		json
//  @Param          comment_id   path      string  true  "comment_id"
//	@Param			request	body		models.CommentUpdateReq	true	"request is required"
//	@Success		200		{object}	models.SuccessResponse{data=models.CommentUpdateRes}
//	@Failure		400		{object}	models.ErrorResponse{errors=interface{}}
//	@Failure		500		{object}	models.ErrorResponse{errors=interface{}}
//	@Security		BearerAuth
//	@Router			/comments/{comment_id} [put]
func (c *CommentControllerImpl) UpdateComment(ctx *gin.Context) {
	var request models.CommentUpdateReq
	commentID := ctx.Param("comment_id")

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
	validateErrs = helpers.CommentUpdateValidator(request)
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

	response, err := c.commentService.UpdateComment(request, userID.(string), commentID)
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
		Message: "Comment updated successfully",
		Data: models.CommentUpdateRes{
			CommentID: response.CommentID,
		},
	})
}

// DeleteComment godoc
//
//	@Summary		delete comment
//	@Description	delete comment
//	@Tags			Comment
//	@Accept			json
//	@Produce		json
//  @Param          comment_id   path      string  true  "comment_id"
//	@Success		200		{object}	models.SuccessResponse{data=models.CommentDeleteRes}
//	@Failure		400		{object}	models.ErrorResponse{errors=interface{}}
//	@Failure		500		{object}	models.ErrorResponse{errors=interface{}}
//	@Security		BearerAuth
//	@Router			/comments/{comment_id} [delete]
func (c *CommentControllerImpl) DeleteComment(ctx *gin.Context) {
	commentID := ctx.Param("comment_id")

	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: "UserID doesn't exist",
		})
		return
	}

	response, err := c.commentService.Delete(commentID, userID.(string))
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
		Message: "Comment deleted successfully",
		Data: models.CommentDeleteRes{
			CommentID: response.CommentID,
		},
	})
}