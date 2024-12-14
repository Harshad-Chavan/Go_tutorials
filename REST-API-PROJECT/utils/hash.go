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

func CheckPassword(password, hashedpassword string) bool {

	// returns nil if passwrod matches the hash else gives error
	err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))

	//hence checking if there is error or not
	return err == nil
}
