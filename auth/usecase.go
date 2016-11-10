package auth

import (
	"errors"

	"fmt"

	"github.com/ReneVallecillo/office.go/domain"
)

//User Entity
type User struct {
	ID       int
	Customer domain.Customer
	Token    string
}

//UserRepository interface to be persisted/retrieved
type UserRepository interface {
	//Save(user User)
	FindByID(id int) (*User, error)
	FindByEmail(email string) (*User, error)
}

//AuthService helps with dependency injection and decoupling
type AuthService struct {
	UserRepository UserRepository
}

//Login2 logs the user
func (auth *AuthService) Login2(email, pass string) (*User, error) {
	fmt.Println("login method reached")
	var authUser *User
	user, err := auth.UserRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if CompareHash(email, pass) {
		authUser, err := auth.UserRepository.FindByID(user.ID)
		if err != nil {
			return nil, err
		}
		authUser.Token = GenerateToken2(*authUser)
	} else {
		err = errors.New("hash not equal")
		return nil, err
	}

	return authUser, nil

}
