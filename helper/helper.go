package helper

import (
	"fmt"
)

// Function in Go can return multiple values
func HandleUserName() (bool, string, string) {
	err := false
	var userName string
	fmt.Println("Enter your first name:")
	// This will not wait for the user to input data rather it executes the next line with nothing in the variable
	// fmt.Scan(userName)
	// This will wait for the input as it has to fill the memory
	fmt.Scan(&userName)

	var lastName string
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	isValidName := len(userName) > 2 && len(lastName) > 2
	if !isValidName {
		fmt.Println("The first and last name should have at least 2 characters each.")
		err = true
	}

	return err, userName, lastName
}
