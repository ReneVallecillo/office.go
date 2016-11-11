package router

import (
	"fmt"
	"net/http"

	"github.com/ReneVallecillo/office.go/domain"
	"github.com/gin-gonic/gin"
)

//LoginRequest used to map request via gin
type LoginRequest struct {
	Email    string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

//AuthService used to decouple code
type AuthService interface {
	Login(email, pass string) (*domain.User, error)
}

//AuthContext gives access to the interface to be used
type AuthContext struct {
	AuthService AuthService
	Authorizer Authorizer
}

//AuthHandler deals with the Auth Request
func (context *AuthContext) AuthHandler(c *gin.Context) {
	var login LoginRequest
	err := c.BindJSON(&login)
	if err != nil {
		fmt.Println(err)
		return
	}
	user, err := context.AuthService.Login(login.Email, login.Password)
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
