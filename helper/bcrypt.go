package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func PassBcrypt(pass string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error bcrypt", err)
		return "", err
	}
	return string(hashed), nil
}