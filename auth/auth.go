package auth

import "github.com/gin-gonic/gin"
import "net/http"

//Login asks for user/pass and validates
//TODO: real thing
func Login(c *gin.Context) {
	content := gin.H{"Hello": "World"}
	c.JSON(http.StatusOK, content)

}
