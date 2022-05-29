package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var availableTickets uint = 50

	greetUser(conferenceName, conferenceTickets, availableTickets)

	// var booking [50]string
	booking := []string{}
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicket := tickets > 0 && availableTickets >= tickets

		if isValidName && isValidEmail && isValidTicket {
			availableTickets = availableTickets - tickets
			booking = append(booking, firstName+" "+lastName)

			fmt.Printf("Thank you %v   %v for booking %v tickets. You will receive email confirmation at %v \n", firstName, lastName, tickets, email)
			fmt.Printf("conferenceName type is %T and conferenceTickets type is %T  \n", conferenceName, conferenceTickets)
			fmt.Printf("\n %v Tickets remains for %v \n", tickets, conferenceName)

			fmt.Printf("\nThe whole slice is %v", booking)
			fmt.Printf("\nThe first slice is %v", booking[0])
			fmt.Printf("\nThe type is %T", booking)
			fmt.Printf("\nThe size of slice is %v \n", len(booking))
			firstNames := []string{}
			for _, bookings := range booking {
				names := strings.Fields(bookings)
				firstNames = append(firstNames, names[0])

			}
			fmt.Printf("\nThese are all first names %v \n", firstNames)

			if availableTickets == 0 {
				fmt.Println("Our conference tickect is fully. Comeback next year")
				break
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
}
func greetUser(conf_Name string, con_tickets int, ava_Tickets uint) {
	fmt.Printf("\n Welcome %v to our conference ticket App\n", conf_Name)
	fmt.Println("Welcome to", conf_Name, "booking application")
	fmt.Printf("We have total of %v tickets and %v are still available", con_tickets, ava_Tickets)
	fmt.Println("Get your ticket here to attend")
}
