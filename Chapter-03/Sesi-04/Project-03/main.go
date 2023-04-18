package main

import (
	"log"
	"project-myGram/database"
	"project-myGram/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	//_ "project-myGram/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const PORT = ":8080"

// @title					MyGram API
// @version					1.0
// @description				This is a MyGram API.
// @host 					localhost:8080
// @BasePath 				/
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	router := gin.Default()

	database.StartDB()
	db := database.GetDB()

	routers.SetupUserRoute(router, db)
	routers.SetupPhotoRoute(router, db)
	routers.SetupSocialRoute(router, db)
	routers.SetupCommentRoute(router, db)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(PORT)
}