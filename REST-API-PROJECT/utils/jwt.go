package utils

import (
	"errors"
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

func VeritfyToken(token string) (error, int64) {

	// this Parse will parse the token
	// post parsing the token it will verify that the method used to sign the token is as expected
	parsedtoken, err := jwt.Parse(token, verifySigningMethod)

	if err != nil {
		return errors.New("could not parse token "), 0
	}

	if !parsedtoken.Valid {
		return errors.New("Invalid Token"), 0
	}

	claims, ok := parsedtoken.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("Invalid claims"), 0
	}

	userid := claims["userid"].(int64)

	return nil, userid
}

func verifySigningMethod(token *jwt.Token) (interface{}, error) {
	// this is to check the type of the signing method
	_, ok := token.Method.(*jwt.SigningMethodHMAC)

	if !ok {
		return nil, errors.New("Unexpected sigining method")
	}
	return []byte(secretkey), nil
}
