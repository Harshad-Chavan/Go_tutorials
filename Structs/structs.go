package main

import (
	"fmt"
	"time"
)

// defining a struct
type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// fucntion attahced to a user now called a method
// (userData user) is called as receiver
func (userData user) outputUserDetails() {
	fmt.Println(userData.firstName, userData.lastName, userData.birthDate)

}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// declaring the varibale as a User type
	var appUser user

	//instantiating the appUser struct
	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
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

	appUser.outputUserDetails()
}

//uisng pointers
// func outputUserDetails(userData *user) {
// 	fmt.Println(userData.firstName, userData.lastName, userData.birthDate)

// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
