package main

import (
	"fmt"
	"strings"
)

func main() {
	// Only applies to variable not constants, also cannot define the type for this syntax
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50 // Only 0 and +ve integers

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	// we can define an array by giving its size and the elemnts data tyoe that it will contain
	// In this case, have the bookings array size as 50 and all of the elements will be of string datatype inside the array
	// var bookings = [conferenceTickets]string{}
	// We can also directly insert the elements like:
	// var bookings = [conferenceTickets]string{"tom", "john"}
	// declaring arrays
	// var bookings [conferenceTickets]string
	// the issue with array is, it occupies the memory as per the memory size for all the elements even if some of the memory does not contain elements
	// this leads to poor memory management as those memory could have been useful for some other things
	// hence we move towards something called SLICE, which is same as array but with dynamic memory allocation
	var bookings []string
	// other ways of declaring and definign slices are:
	// var bookings = []string{}
	// bookings := []string{}

	// for {}: is a inifinite loop
	for {
		invalidName, userName, lastName := handleUserName()
		if invalidName {
			continue
		}

		invalidEmail, email := handleEmail()
		if invalidEmail {
			continue
		}

		invalidTickets, userTickets := handleTickets(remainingTickets)
		if invalidTickets {
			continue
		}

		bookTickets(remainingTickets, bookings, userName, lastName, email, conferenceName, userTickets)

		firstNames := getFirstNames(bookings)
		fmt.Printf("See your name in our booking list: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("our conference is booked out. Come back next year.")
			break
		}
	}
}

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Println("We have total of", confTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	// range: this makes the slice iterable and returns 2 values: index and the item/element of each slice at that index of the slice
	// But, we do not need the index for this use case, hence, we can substitute the index with an _
	// this states to the Go compiler that we do not need anything as index as we don't seem to use it right now
	for _, booking := range bookings {
		// Fields: this is a method of hte string strings package which splits the the string by spaces
		// this is similar to string,split(' ') in javascript, note the space inside rhe split
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

// Function in Go can return multiple values
func handleUserName() (bool, string, string) {
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

func handleEmail() (bool, string) {
	err := false
	var email string
	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	isValidEmail := strings.Contains(email, "@")
	if !isValidEmail {
		fmt.Println("Please, enter a valid email address")
		err = true
	}
	return err, email
}

func handleTickets(remainingTickets uint) (bool, int) {
	err := false
	var userTickets int
	fmt.Println("Enter the number of tickets:")
	fmt.Scan(&userTickets)

	if userTickets < 1 {
		fmt.Println("Please provide a valid number")
		err = true
		return err, userTickets
	}

	if uint(userTickets) > remainingTickets {
		fmt.Printf("We only have %v remaining tickets to book. Please, book accordingly\n", remainingTickets)
		err = true
	}

	return err, userTickets
}

func bookTickets(remainingTickets uint, bookings []string, userName string, lastName string, email string, conferenceName string, userTickets int) {
	// As the userTickets is an integer, but remainingTickets is uint, hence we need to convert either one of them to other type
	remainingTickets = remainingTickets - uint(userTickets)
	bookings = append(bookings, userName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
