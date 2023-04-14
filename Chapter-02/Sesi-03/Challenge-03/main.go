package main

import (
	"simple-rest-api-with-db/database"
	routers "simple-rest-api-with-db/routers"
	_ "github.com/lib/pq"
)

func main() {
	db := database.DbConnection()
	defer db.Close()

	const PORT = ":8080"

	routers.StartServer().Run(PORT)
}



