package handlers

import (
	"net/http"

	"github.com/ReneVallecillo/office.go/model"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//UserListHandler responds with a user list json
func UserListHandler(c *gin.Context) {

	db := c.MustGet("DB").(*sqlx.DB)
	user := model.User{}
	users, err := user.UserList(db)

	if err != nil {
		errorResponse := gin.H{"content": err}
		c.JSON(http.StatusNotFound, errorResponse)
		return
	}

	content := gin.H{"content": users}
	c.JSON(http.StatusOK, content)

}
