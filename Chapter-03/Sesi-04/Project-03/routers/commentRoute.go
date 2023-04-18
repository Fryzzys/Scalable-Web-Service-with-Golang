package routers

import (
	"project-myGram/controllers"
	"project-myGram/middlewares"
	"project-myGram/repository"
	"project-myGram/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupCommentRoute(router *gin.Engine, db *gorm.DB) {
	commentRepository := repository.NewCommentRepository(db)
	photoRepository := repository.NewPhotoRepository(db)
	commentService := services.NewCommentService(commentRepository, photoRepository)
	commentController := controllers.NewCommentController(commentService)

	authUser := router.Group("/comments", middlewares.AuthMiddleware)
	{
		authUser.POST("/:photo_id", commentController.CreateComment)
		authUser.GET("", commentController.GetAll)
		authUser.GET("/:comment_id", commentController.GetOne)
		authUser.PUT("/:comment_id", commentController.UpdateComment)
		authUser.DELETE("/:comment_id", commentController.DeleteComment)
	}
}