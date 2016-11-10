package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Database is a middleware that add a db cnx to the context
func Database(db *sqlx.DB) gin.HandlerFunc {
	fmt.Println("adding db to context")
	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}
