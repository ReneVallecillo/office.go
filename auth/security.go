package auth

import (
	"encoding/base64"
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

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
