package routers

import (
	"project-myGram/controllers"
	"project-myGram/middlewares"
	"project-myGram/repository"
	"project-myGram/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoute(router *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	photoRepository := repository.NewPhotoRepository(db)
	socialRepository := repository.NewSocialRepository(db)
	userService := services.NewUserService(userRepository, photoRepository, socialRepository)
	userController := controllers.NewUserController(userService)

	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)

	authUser := router.Group("/user", middlewares.AuthMiddleware)
	{
		authUser.GET("/profile", userController.GetProfile)
	}
}