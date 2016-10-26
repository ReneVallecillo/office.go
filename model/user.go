package model

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// User holds the mapping of the user object
type User struct {
	ID            uint32 `db:"user_id" json:"id"` // Don't use Id, use UserID() instead for consistency with MongoDB
	FirstName     string `db:"first_name" json:"first_name"`
	LastName      string `db:"last_name" json:"last_name"`
	Email         string `db:"email" json:"email"`
	Password      string `db:"password" json:"password"`
	StatusID      uint8  `db:"status_id" json:"status_id"`
	Address       string `db:"address" json:"address"`
	ContactNumber string `db:"contact_number" json:"contact_number"`
	GenderID      string `db:"gender_id" json:"gender_id"`
	PicURL        string `db:"pic_url" json:"pic_url"`
	UserLevel     string `db:"user_level" json:"user_level"`

	Control
	Token string `json:"token"`
}

// UserFindByID returns one User given the ID
func (user *User) UserFindByID(db *sqlx.DB, id int) (User, error) {
	start := time.Now()
	result := User{}
	query := `SELECT user_id, first_name, last_name, email 
              FROM "user" WHERE user_id = $1;`
	err := db.Get(&result, query, id)
	elapsed := time.Since(start)
	log.Printf("DB request took %s", elapsed)
	return result, err
}

//UserList list all users
func (user *User) UserList(db *sqlx.DB) ([]User, error) {
	users := []User{}
	query := `SELECT * FROM "user"`
	err := db.Select(&users, query)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to query for users")
	}
	return users, nil
}
