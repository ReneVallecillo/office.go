package postgres

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

//TODO: USE ENV IN PRODUCTION
const (
	DB_USER     = "OfficeAdmin"
	DB_PASSWORD = "office123"
	DB_NAME     = "Office_Dev"
	DB_PORT     = "5432"
	DB_HOST     = "localhost"
)

// InitDB initializes the DB
func InitDB() *sqlx.DB {
	// TODO: Use config files
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME, DB_HOST, DB_PORT)
	db, err := sqlx.Connect("postgres", dbinfo)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
