package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// defining a struct
type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// fucntion to create a struct
// adding some validation
func newUser(userFirstName, userLastName, userBirthdate string) (*user, error) {
	if userFirstName == "" || userLastName == "" || userBirthdate == "" {
		return nil, errors.New("all values are required")
	}

	return &user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}, nil
}

// fucntion attahced to a user now called a method
// (userData user) is called as receiver
func (userData user) outputUserDetails() {
	fmt.Println(userData.firstName, userData.lastName, userData.birthDate)

}

func (userData *user) clearUserName() {
	userData.firstName = ""
	userData.lastName = ""

}

func (userData *user) capitalize() {
	userData.firstName = strings.ToUpper(userData.firstName)
	userData.lastName = strings.ToUpper(userData.lastName)

}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// declaring the varibale as a User type
	var appUser *user
	var err error

	//instantiating the appUser struct
	appUser, err = newUser(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(appUser)

	// alternate way of instantiating make sure the order is same
	// appUser = user{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthdate,
	// 	time.Now(),
	// }

	//we can create an empty struct also
	// var appUser_emp user = user{}
	// fmt.Println(appUser_emp)

	// ... do something awesome with that gathered data!

	// outputUserDetails(&appUser)
	appUser.capitalize()
	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

//uisng pointers
// func outputUserDetails(userData *user) {
// 	fmt.Println(userData.firstName, userData.lastName, userData.birthDate)

// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
