package main

import (
	database "github.com/ReneVallecillo/office.go/postgres"
	"github.com/ReneVallecillo/office.go/router"

	_ "github.com/lib/pq"
)

func main() {
	db := database.InitDB()
	defer db.Close()

	router := router.InitRouter(db)
	//adding db middleware

	router.Run(":8080")
}
