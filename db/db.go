package db

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

// Database is a middleware that add a db cnx to the context
func Database(db *sqlx.DB) gin.HandlerFunc {
	fmt.Println("adding db to context")
	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}
