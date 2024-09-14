package user

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// defining a struct
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string, user *User) Admin {
	return Admin{
		email:    email,
		password: password,
		User:     *user,
	}

}

// fucntion to create a struct
// adding some validation
func NewUser(userFirstName, userLastName, userBirthdate string) (*User, error) {
	if userFirstName == "" || userLastName == "" || userBirthdate == "" {
		return nil, errors.New("all values are required")
	}

	return &User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}, nil
}

// fucntion attahced to a user now called a method
// (userData user) is called as receiver
func (userData User) OutputUserDetails() {
	fmt.Println(userData.firstName, userData.lastName, userData.birthDate)

}

func (userData *User) ClearUserName() {
	userData.firstName = ""
	userData.lastName = ""

}

func (userData *User) Capitalize() {
	userData.firstName = strings.ToUpper(userData.firstName)
	userData.lastName = strings.ToUpper(userData.lastName)

}
