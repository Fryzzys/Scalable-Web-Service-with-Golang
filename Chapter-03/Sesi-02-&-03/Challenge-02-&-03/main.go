package main

import (
	"simple-rest-api-with-jwt/database"
	"simple-rest-api-with-jwt/routers"
	"log"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	database.StartDB()
	r := routers.StartApp()
	r.Run(":8080")
}
