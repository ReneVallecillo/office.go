package auth

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"bitbucket.org/reneval/lawparser/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Username string `json:"username"`
	// recommended having
	jwt.StandardClaims
}

//HashPass takes a pass (string) and returns a base64 hash
func HashPass(pass string) (string, error) {

	// TODO: use config files
	var cost = 1
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	if err != nil {
		return "", errors.Wrap(err, "Could not generate pass hash")
	}

	// Encode the hash as base64 and return
	hashBase64 := base64.StdEncoding.EncodeToString(hash)

	return hashBase64, nil

}

// CompareHash comares 2 hashes passwords
func CompareHash(reqPass, dbPass string) bool {
	//Decode
	hashBytes, err := base64.StdEncoding.DecodeString(dbPass)
	if err != nil {
		err := errors.Wrap(err, "Invalid base64 string")
		fmt.Println(err)
		return false
	}

	err = bcrypt.CompareHashAndPassword(hashBytes, []byte(reqPass))
	return err == nil

}

// GenerateToken generates a jwt token
func GenerateToken(user *models.User) string {
	// Expires the token and cookie in 1 hour
	expireToken := time.Now().Add(time.Hour * 24).Unix()
	//expireCookie := time.Now().Add(time.Hour * 1)

	// We'll manually assign the claims but in production you'd insert values from a database
	claims := Claims{
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "localhost:9000", //TODO: use real info
		},
	}

	// Create the token using your claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signs the token with a secret.
	//TODO: USE ENV for secret
	signedToken, _ := token.SignedString([]byte("secret"))

	return signedToken
}

// Middleware to protect /profile and /logout
func validate(protectedPage http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		//Validate the token and if it passes call the protected handler below.
		protectedPage(res, req)
	})
}
