package routers

import (
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/controllers"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/middlewares"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/repository"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/services"
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