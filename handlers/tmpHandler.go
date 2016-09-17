package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NotImplemented is a temp function to fill Non Implemented handlers
func NotImplemented(c *gin.Context) {
	content := gin.H{"Response": "Not Implemented"}
	c.JSON(http.StatusOK, content)
}
