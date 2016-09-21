package model

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

// User holds the mapping of the user object
type User struct {
	ID        uint32 `db:"user_id" json:"id"` // Don't use Id, use UserID() instead for consistency with MongoDB
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
	StatusID  uint8  `db:"status_id" json:"status_id"`
	Control
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
