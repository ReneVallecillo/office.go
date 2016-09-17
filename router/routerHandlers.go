package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping checks if api is live
func Ping(c *gin.Context) {
	content := gin.H{"Content": "Api is up and running"}
	c.JSON(http.StatusOK, content)
}
