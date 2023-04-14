package database

import (
	"fmt"
	"database/sql"
	
)

var (
	db  *sql.DB
	err error
)

const (
	HOST	 = "localhost"
	PORT	 = 5432
	USERNAME = "postgres"
	PASSWORD = "1"
	DBNAME	 = "library"
)

func DbConnection() *sql.DB {

	sqldat := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
				HOST, PORT, USERNAME, PASSWORD, DBNAME)
	db, err = sql.Open("postgres", sqldat)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	return db
}
