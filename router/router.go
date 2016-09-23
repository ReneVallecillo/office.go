package router

import (
	"github.com/ReneVallecillo/office.go/auth"
	database "github.com/ReneVallecillo/office.go/db"
	"github.com/ReneVallecillo/office.go/handlers"
	"github.com/ReneVallecillo/office.go/mock"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// InitRouter initializes the router
func InitRouter(db *sqlx.DB) *gin.Engine {
	router := gin.Default()
	router.Use(database.Database(db))

	v1 := router.Group("/api/v1")
	{
		//Routes
		v1.POST("/login", auth.Login)
		v1.GET("/", handlers.NotImplemented)
		v1.GET("/ping", Ping)
	}

	//Mock Routes
	test := router.Group("/test")
	{
		test.GET("/products/:slug/find", mock.MockHandler)
		test.GET("/products", mock.MockProductHandler)
	}

	authorized := router.Group("/", auth.TokenAuthMiddleware())
	{
		authorized.GET("/profile", mock.MockProductHandler)
	}

	return router
}
