package helper

import (
	"log"
	"time"
	"yuhuuuMit/config"

	"github.com/golang-jwt/jwt"
)

func TokenGenerate(id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorizad"] = true
	claims["userID"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // token expired
	tokenHash := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenHash.SignedString([]byte(config.JWTKey))
	if err != nil {
		log.Println("error generate token jwt ~", err)
		return "", err
	}
	return token, nil
}