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
	Login2(email, pass string) (*domain.User, error)
}

//AuthContext gives access to the interface to be used
type AuthContext struct {
	AuthService AuthService
}

//AuthHandler deals with the Auth Request
func (context *AuthContext) AuthHandler(c *gin.Context) {
	fmt.Println("Handler Reached")
	var login LoginRequest
	err := c.BindJSON(&login)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("before login")
	user, err := context.AuthService.Login2(login.Email, login.Password)
	if err != nil {
		//err = errors.Wrap(err, "couldn't find user")
		c.JSON(http.StatusOK, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
