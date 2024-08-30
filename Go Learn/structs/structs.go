package main

import (
	"fmt"

	"examples.com/struct/user"
)

func main() {
	userFirstName := getUserData("Please Enter your first name: ")
	userLastName := getUserData("Please Enter your last name: ")
	userBirthDate := getUserData("Please enter your dob in MM/DD/YYYY: ")

	var appUser *user.User

	// appUser, err := newUser(userFirstName, userLastName, userBirthDate)
	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("admin@test.com", "admin@omz.com")
	
	admin.User.OutputUserDetails()
	admin.User.ClearUserName()
	admin.User.OutputUserDetails()

	appUser.OutputUserDetails() //don't want to pass any argument.
	appUser.ClearUserName()
	appUser.OutputUserDetails()
	// outputUserDetails(&appUser)
	// fmt.Println(firstName, lastName, birthDate)
}

// func outputUserDetails(firstName, lastName, bithDate string) {
// 	//...
// 	fmt.Println(firstName, lastName, bithDate)
// }

// func outputUserDetails(u *user) {
// 	//...
// 	fmt.Println(u.firstName, u.lastName, u.birthDate)
// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
