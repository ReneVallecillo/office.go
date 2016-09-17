package router

import (
	"github.com/ReneVallecillo/office/auth"
	"github.com/ReneVallecillo/office/handlers"
	"github.com/ReneVallecillo/office/mock"
	"github.com/gin-gonic/gin"
)

// InitRouter initializes the router
func InitRouter() *gin.Engine {
	router := gin.Default()

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

	return router

}
