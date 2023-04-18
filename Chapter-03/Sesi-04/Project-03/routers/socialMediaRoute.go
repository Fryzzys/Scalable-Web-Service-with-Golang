package routers

import (
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/controllers"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/middlewares"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/repository"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/services"
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