package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretkey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {

	// Generating token and adding claims to the token
	// setting an expire time
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"expire": time.Now().Add(time.Hour * 2).Unix(),
	})

	// token is generated but we convert it into a string with a secret key
	return token.SignedString([]byte(secretkey))
}
