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
	var adminUserDets *user.User
	var err error

	//instantiating the appUser struct
	appUser, err = user.NewUser(userFirstName, userLastName, userBirthdate)
	adminUserDets, err = user.NewUser("admin_harshad", "chavan", "05/12/12")

	if err != nil {
		fmt.Println(err)
		return
	}

	adminUser := user.NewAdmin("test@ex.com", "test_123", adminUserDets)

	// fmt.Println(appUser)
	// fmt.Println(adminUser)

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

	// this if you give a name to the embeded structure
	// adminUser.User.Capitalize()
	// adminUser.User.OutputUserDetails()

	adminUser.Capitalize()
	adminUser.OutputUserDetails()
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
