package routers

import (
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/controllers"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/middlewares"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/repository"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/services"
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