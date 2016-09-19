package main

import (
	database "github.com/ReneVallecillo/office.go/db"
	"github.com/ReneVallecillo/office.go/router"
)

func main() {
	db := database.InitDB()
	defer db.Close()

	router := router.InitRouter(db)
	//adding db middleware

	router.Run(":8080")
}
