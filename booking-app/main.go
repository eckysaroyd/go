package main

import (
	"booking-app/helper"
	"fmt"
)

var conferenceName = "Go Conference"

const conferenceTickets int = 50

var AvailableTickets uint = 50

// var booking [50]string
// var booking = make([]map[string]string, 0)
var booking = make([]userData, 0)

type userData struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

func main() {

	// function call for greetUser func() before package level variable
	// greetUser(conferenceName, conferenceTickets, availableTickets)

	// function call for greetUser func() after package level variable
	greetUser()
	firstName, lastName, email, tickets := userInput()
	// validation function
	isValidName, isValidEmail, isValidTicket := helper.Validate(firstName, lastName, email, tickets, AvailableTickets)

	if isValidName && isValidEmail && isValidTicket {
		bookTicket(tickets, firstName, lastName, email)

		// function call for first_names func()
		get_firstNames := first_name()
		fmt.Printf("\nThese are all first names %v \n", get_firstNames)

		if AvailableTickets == 0 {
			fmt.Println("Our conference tickect is fully. Comeback next year")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("FirstName or LastName is too short")
		}
		if !isValidEmail {
			fmt.Println("Email is invalid")
		}
		if !isValidTicket {
			fmt.Println("wrong ticke no")
		}
		// fmt.Printf("\nWe only have %v Tickets. You cant book %v tickets\n", availableTickets, tickets)
		// fmt.Println("Your Input is Invalid")

	}

}
func greetUser() {
	fmt.Printf("\n Welcome %v to our conference ticket App\n", conferenceName)
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("We have total of %v tickets and %v are still available", conferenceTickets, AvailableTickets)
	fmt.Println("Get your ticket here to attend")
}
func first_name() []string {
	firstNames := []string{}
	for _, bookings := range booking {
		firstNames = append(firstNames, bookings.firstName)
		// bookings["firstName"] used in map
	}
	return firstNames
}

// func validate(firstName string, lastName string, email string, tickets uint) (bool, bool, bool) {
// 	isValidName := len(firstName) >= 2 && len(lastName) >= 2
// 	isValidEmail := strings.Contains(email, "@")
// 	isValidTicket := tickets > 0 && availableTickets >= tickets

// 	return isValidName, isValidEmail, isValidTicket
// }
func userInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var tickets uint
	for {
		fmt.Printf("Enter Your First Name \n")
		fmt.Scan(&firstName)
		fmt.Printf("Enter Your Last Name \n")
		fmt.Scan(&lastName)
		fmt.Printf("Enter Your Email Name \n")
		fmt.Scan(&email)
		fmt.Printf("Enter Your Tickets_no \n")
		fmt.Scan(&tickets)

		return firstName, lastName, email, tickets
	}
}
func bookTicket(tickets uint, firstName string, lastName string, email string) {

	AvailableTickets = AvailableTickets - tickets

	// create a struct
	var userData = userData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		tickets:   tickets,
	}

	// create a map for a user
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTicket"] = strconv.FormatUint(uint64(tickets), 10)

	booking = append(booking, userData)
	fmt.Printf("The list of Booking is %v \n", booking)

	fmt.Printf("Thank you %v   %v for booking %v tickets. You will receive email confirmation at %v \n", firstName, lastName, tickets, email)
	fmt.Printf("conferenceName type is %T and conferenceTickets type is %T  \n", conferenceName, conferenceTickets)
	fmt.Printf("\n %v Tickets remains for %v \n", AvailableTickets, conferenceName)
	fmt.Printf("\nThe whole slice is %v", booking)
	fmt.Printf("\nThe first slice is %v", booking[0])
	fmt.Printf("\nThe type is %T", booking)
	fmt.Printf("\nThe size of slice is %v \n", len(booking))
}
