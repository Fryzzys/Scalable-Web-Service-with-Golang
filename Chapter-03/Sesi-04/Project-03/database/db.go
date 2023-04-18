package database

import (
	"fmt"
	"log"
	"os"
	"github.com/Fryzzys/Scalable-Web-Service-with-Golang/Chapter-03/Sesi-04/Project-03/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	var (
		DB_HOST     = os.Getenv("DB_HOST")
		DB_USER     = os.Getenv("DB_USER")
		DB_PASSWORD = os.Getenv("DB_PASSWORD")
		DB_NAME     = os.Getenv("DB_NAME")
		DB_PORT     = os.Getenv("DB_PORT")
	)

	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database", err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.SocialMedia{}, models.Comment{})
}

func GetDB() *gorm.DB {
	return db
}