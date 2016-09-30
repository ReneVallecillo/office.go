package auth

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/ReneVallecillo/office.go/model"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// LoginUser is a tmp struct that hold minimal data to auth user
type LoginUser struct {
	ID       int            `db:"user_id"`
	Password sql.NullString `db:"password"`
}

type LoginRequest struct {
	Email    string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

//Login asks for user/pass and validates
//TODO: Add jwt logic
func Login(c *gin.Context) {
	query := `SELECT user_id, password FROM "user" WHERE "email" = $1`
	db := c.MustGet("DB").(*sqlx.DB)
	var login LoginRequest
	err := c.BindJSON(&login)
	if err != nil {
		fmt.Println(err)
		return
	}

	// email := c.PostForm("email")
	// pass := c.PostForm("password")
	// fmt.Println("pass:", pass)

	loginUser := LoginUser{}
	err = db.Get(&loginUser, query, login.Email)
	if err != nil {
		err = errors.Wrap(err, "couldn't find user")
		c.JSON(http.StatusOK, err.Error())
		return
	}

	if CompareHash(login.Password, loginUser.Password.String) {
		user := model.User{}
		user, err := user.UserFindByID(db, loginUser.ID)
		if err != nil {
			err = errors.Wrap(err, "DB error")
			c.JSON(http.StatusOK, err.Error())
			return
		}

		//Set token
		token := GenerateToken(user)
		user.Token = token
		SetSession(c, "Auth", token)
		c.JSON(http.StatusOK, user)

	} else {
		err = errors.New("Pass mismatch")
		c.JSON(http.StatusOK, err.Error())
		return
	}

}

func SetSession(c *gin.Context, name string, token string) {
	c.SetCookie(
		name,
		token,
		1000,
		"/",
		"",
		true,
		true)
}

func Logout(c *gin.Context) {
	c.SetCookie(
		"Auth",
		"none",
		-1,
		"/",
		"localhost",
		true,
		true,
	)
}
