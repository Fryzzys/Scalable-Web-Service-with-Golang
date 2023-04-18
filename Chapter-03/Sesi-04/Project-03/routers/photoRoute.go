package routers

import (
	"project-myGram/controllers"
	"project-myGram/middlewares"
	"project-myGram/repository"
	"project-myGram/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupPhotoRoute(router *gin.Engine, db *gorm.DB) {
	photoRepository := repository.NewPhotoRepository(db)
	commentRepository := repository.NewCommentRepository(db)
	photoService := services.NewPhotoService(photoRepository, commentRepository)
	photoController := controllers.NewPhotoController(photoService)

	authUser := router.Group("/photos", middlewares.AuthMiddleware)
	{
		authUser.POST("", photoController.CreatePhoto)
		authUser.GET("", photoController.GetAll)
		authUser.GET("/:photo_id", photoController.GetOne)
		authUser.PUT("/:photo_id", photoController.UpdatePhoto)
		authUser.DELETE("/:photo_id", photoController.DeletePhoto)
	}
}