package postgres

import (
	"fmt"

	"github.com/ReneVallecillo/office.go/domain"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type UserService struct {
	DB *sqlx.DB
}

//FindByID return one user with the query ID
func (db *UserService) FindByID(id uint32) (*domain.User, error) {
	var user = domain.User{}

	query := `SELECT user_id, first_name, last_name, email 
              FROM "user" WHERE user_id = $1;`
	err := db.DB.Get(&user, query, id)
	if err != nil {
		err = errors.Wrap(err, "couldn't find user by id")
		return nil, err
	}
	return &user, nil

}

//FindByEmail finds an User by his email
func (db *UserService) FindByEmail(email string) (*domain.User, error) {
	fmt.Println("llego a postgres")
	var user domain.User
	query := `SELECT user_id, password FROM "user" WHERE "email" = $1`
	err := db.DB.Get(&user, query, email)
	if err != nil {
		err = errors.Wrap(err, "couldn't find user by email")
		return nil, err
	}

	return &user, nil

}

//FindAll returns all users in DB.
func (db *UserService) FindAll() ([]*domain.User, error) {
	fmt.Println("llego a postgres")
	var users []*domain.User
	query := `SELECT * FROM "user"`
	err := db.DB.Select(&users, query)
	if err != nil {
		err = errors.Wrap(err, "couldn't find any user")
		return nil, err
	}

	return users, nil

}
