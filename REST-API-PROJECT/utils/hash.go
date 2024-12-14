package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	fmt.Println(password)
	bytespassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytespassword), err

}
