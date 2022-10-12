package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
	"sync"
	"time"
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
// var bookings = make([]map[string]string, 0) ---> OLD VERSION CODE

var bookings = make([]UserData, 0)

// Creating a structure which can contain data of different types unlike map
// It is basically saying to create a custom data type according to the user's need
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

// other ways of declaring and definign slices are:
// var bookings = []string{}
// bookings := []string{}

// If we remove the for loop, then the main thread will stop when 1 ticket is booked, the program finishes
// The problem with creating goroutine is, this will stop execution if the main thread stops/finishes
// To tackle that we need to let the main thread know that we need to wait as some background process is going on
// We need something called WaitGroup to do that
var wg = sync.WaitGroup{}

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

		bookedUser := bookTickets(userName, lastName, email, userTickets)
		// This will inform the main thread that it needs to wait for 1 task to finish
		// After this we can write the "**go**" keyword to make another goroutine
		// Similarly, we can let the main thread know how much goroutines to wait and then create them accordingly
		wg.Add(1)
		// the "**go**" keyword makes the block of code asynchronous by assigning it to another thread called goroutine
		// The new thread runs in the background while the main thread conitnue to execute
		// At some point of time the background thread will finish the execution and its output will be printed in the console
		go sendTicket(bookedUser)

		firstNames := getFirstNames()
		fmt.Printf("See your name in our booking list: %v\n", firstNames)
		break

		// if remainingTickets == 0 {
		// 	fmt.Println("our conference is booked out. Come back next year.")
		// 	break
		// }
	}
	// This will make the main thread to wait till all the goroutines is finished
	wg.Wait()
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

		firstNames = append(firstNames, booking.firstName)
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

func bookTickets(userName string, lastName string, email string, userTickets int) UserData {

	// Created a map with the key as a string and the value as a string
	userData := UserData{
		firstName:       userName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// As the userTickets is an integer, but remainingTickets is uint, hence we need to convert either one of them to other type
	remainingTickets = remainingTickets - uint(userTickets)
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	return userData
}

func sendTicket(user UserData) {
	// Waits for 5 seconds to simulate asynchronous nature of creating and sending a ticket to email
	time.Sleep(5 * time.Second)
	// Sprintf() will return the string that would print in the terminal rather than printing it.
	// The returned string is kept in variable for later usage
	ticket := fmt.Sprintf("%v tickets for %v %v", user.numberOfTickets, user.firstName, user.lastName)
	fmt.Println("------ TICKET ------")
	fmt.Printf("Sending ticket...\n %v \nto email address %v\n", ticket, user.email)
	fmt.Println("------ TICKET ------")

	// This will signal the waitgroup stating that thi goroutine is finished
	wg.Done()
}
