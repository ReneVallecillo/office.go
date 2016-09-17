package router

import (
	"github.com/ReneVallecillo/office/auth"
	"github.com/gin-gonic/gin"
)

// InitRouter initializes the router
func InitRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		//Routes
		v1.POST("/login", auth.Login)
	}

	return router

}
