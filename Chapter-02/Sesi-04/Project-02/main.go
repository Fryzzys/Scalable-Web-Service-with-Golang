package main

import (
	"project-gorm/database"
	"project-gorm/routes"
	"github.com/gin-gonic/gin"
)

const PORT = ":8080"

func main() {
	router := gin.Default()

	database.StartDB()
	db := database.GetDB()

	routes.SetupBookRoute(router, db)

	router.Run(PORT)
}