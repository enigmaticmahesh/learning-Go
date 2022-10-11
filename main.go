package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
	"strings"
)

// we cannot use := in package level variables
var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50 // Only 0 and +ve integers
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
// var bookings []string --> OLD VERSION CODE

// Created list of maps for storing userData(user object as in javascript)
// For creating list of data using make(), we need to pass the length as param
// Here, we just made the list with length 0, as the data added it will increase the size similar to slice
var bookings = make([]map[string]string, 0)

// other ways of declaring and definign slices are:
// var bookings = []string{}
// bookings := []string{}

func main() {

	greetUsers()

	// for {}: is a inifinite loop
	for {
		invalidName, userName, lastName := helper.HandleUserName()
		if invalidName {
			continue
		}

		invalidEmail, email := handleEmail()
		if invalidEmail {
			continue
		}

		invalidTickets, userTickets := handleTickets()
		if invalidTickets {
			continue
		}

		bookTickets(userName, lastName, email, userTickets)

		firstNames := getFirstNames()
		fmt.Printf("See your name in our booking list: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("our conference is booked out. Come back next year.")
			break
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	// range: this makes the slice iterable and returns 2 values: index and the item/element of each slice at that index of the slice
	// But, we do not need the index for this use case, hence, we can substitute the index with an _
	// this states to the Go compiler that we do not need anything as index as we don't seem to use it right now
	for _, booking := range bookings {
		// Fields: this is a method of hte string strings package which splits the the string by spaces
		// this is similar to string,split(' ') in javascript, note the space inside rhe split
		// var names = strings.Fields(booking) ---> OLD VERSION CODE

		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
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

func handleTickets() (bool, int) {
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

func bookTickets(userName string, lastName string, email string, userTickets int) {

	// Created a map with the key as a string and the value as a string
	userData := make(map[string]string)
	userData["firstName"] = userName
	userData["lastName"] = lastName
	userData["email"] = email
	// Converting "int" data to string using the "strconv" package available by Go
	// "10" is used to get the number in decimal, other can be "16" for hexadecimal
	userData["userTickets"] = strconv.FormatInt(int64(userTickets), 10)

	// As the userTickets is an integer, but remainingTickets is uint, hence we need to convert either one of them to other type
	remainingTickets = remainingTickets - uint(userTickets)
	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
