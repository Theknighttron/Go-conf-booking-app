package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName string = "Go Conference"
	var remainingTickets uint = 50
	const conferenceTickets = 50
	var bookings []string

	// show welcome message
	greetUser(conferenceName, conferenceTickets, remainingTickets)

	for {

		// get user input
		firstName, lastName, email, userTickets := getUserInput()

		// check the length of first and last name
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			// booking logic
			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)

			// iterate over names and return only the firstnames
			firstNames := getFirstNames(bookings)
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

func greetUser(confName string, confTickets int, availableTickets uint) {
	fmt.Println("Welcome to", confName, "booking application")
	fmt.Println("We have total of", confTickets, "tickets and", availableTickets, "are still available")
	fmt.Println("Get your ticket here to attend")
}

func getFirstNames(bookings []string) []string {

	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
