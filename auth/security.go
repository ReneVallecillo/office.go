package auth

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/ReneVallecillo/office.go/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
func GenerateToken(user model.User) string {
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

// TokenAuthMiddleware exists to protect /profile and /logout
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("Auth")
		if err != nil {
			fmt.Println(err)
			RespondWithError(http.StatusUnauthorized, "Token Required", c)
			return
		}

		//Return token
		token, err := jwt.ParseWithClaims(
			cookie,
			&Claims{},
			func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method")
				}
				return []byte("secret"), nil
			})

		if err != nil {
			detail := errors.Wrap(err, "Token Invalid")
			fmt.Printf("%v", detail)
			RespondWithError(http.StatusUnauthorized, "Token Invalid", c)
			return
		}

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			c.Set("claim", *claims)
		} else {
			c.JSON(http.StatusUnauthorized, "Invalid Token")
			return
		}

		c.Next()
	}
}

// RespondWithError ends up the request chain.
func RespondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	c.Abort()
}
