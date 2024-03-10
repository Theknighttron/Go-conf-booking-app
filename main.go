package main

import (
	"fmt"
	"strings"
)

const conferenceTickets = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	// show welcome message
	greetUser()

	for {

		// get user input
		firstName, lastName, email, userTickets := getUserInput()

		// check the length of first and last name
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			// booking logic
			bookTicket(userTickets, firstName, lastName, email)

			// iterate over names and return only the firstnames
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			fmt.Printf("These are all our bookings: %v\n", bookings)
			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				// end program
				fmt.Println("Our conference is booked up, Come back next year.")
				break
			}
		} else {
			fmt.Print("\n")
			if !isValidName {
				fmt.Println("> Firstname or lastname is too short")
			}
			if !isValidEmail {
				fmt.Println("> Email address does not contain @ symbol")
			}
			if !isValidTicketNumber {
				fmt.Println("Invalid number of tickets")
			}
			fmt.Printf("> Please, try again...\n")
		}

	}

}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("\n")
	fmt.Print("Enter your Firstname: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your Lastname: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your Email: ")
	fmt.Scan(&email)

	fmt.Print("Enter your Number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func greetUser() {
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your ticket here to attend")
}

func getFirstNames() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
