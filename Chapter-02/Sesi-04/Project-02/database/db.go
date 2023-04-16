package database

import (
	"fmt"
	"log"
	"project-gorm/models/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "1"
	dbName   = "library"
	port     = 5432
	db       *gorm.DB
	err      error
)

func StartDB() {
	connectionConfig := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, password, dbName, port)

	db, err = gorm.Open(postgres.Open(connectionConfig), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database", err)
	}

	db.Debug().AutoMigrate(entity.Book{})
}

func GetDB() *gorm.DB {
	return db
}