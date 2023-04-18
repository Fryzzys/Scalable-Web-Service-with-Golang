package routers

import (
	"project-myGram/controllers"
	"project-myGram/middlewares"
	"project-myGram/repository"
	"project-myGram/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupSocialRoute(router *gin.Engine, db *gorm.DB) {
	socialRepository := repository.NewSocialRepository(db)
	socialService := services.NewSocialService(socialRepository)
	socialController := controllers.NewSocialController(socialService)

	authUser := router.Group("/social-media", middlewares.AuthMiddleware)
	{
		authUser.POST("", socialController.CreateSocial)
		authUser.GET("", socialController.GetAll)
		authUser.GET("/:social_media_id", socialController.GetOne)
		authUser.PUT("/:social_media_id", socialController.UpdateSocialMedia)
		authUser.DELETE("/:social_media_id", socialController.DeleteSocialMedia)
	}
}