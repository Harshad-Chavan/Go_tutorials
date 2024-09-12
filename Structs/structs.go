package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// declaring the varibale as a User type
	var appUser *user.User
	var err error

	//instantiating the appUser struct
	appUser, err = user.NewUser(userFirstName, userLastName, userBirthdate)

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
	appUser.Capitalize()
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
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
