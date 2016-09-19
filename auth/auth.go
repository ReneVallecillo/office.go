package auth

import (
	"database/sql"
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

//Login asks for user/pass and validates
//TODO: Add jwt logic
func Login(c *gin.Context) {
	query := `SELECT "user_id", "password" FROM "user" WHERE "email" = $1`
	db := c.MustGet("DB").(*sqlx.DB)

	email := c.PostForm("email")
	pass := c.PostForm("pass")

	loginUser := LoginUser{}
	err := db.Get(&loginUser, query, email)
	if err != nil {
		err = errors.Wrap(err, "couldn't find user")
		c.JSON(http.StatusOK, err.Error())
		return
	}

	if CompareHash(pass, loginUser.Password.String) {
		user := model.User{}
		user, err := user.UserFindByID(db, loginUser.ID)
		if err != nil {
			err = errors.Wrap(err, "DB error")
			c.JSON(http.StatusOK, err.Error())
			return
		}
		c.JSON(http.StatusOK, user)

	} else {
		err = errors.New("Pass mismatch")
		c.JSON(http.StatusOK, err.Error())
		return
	}

}
