package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var availableTickets uint = 50
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("We have total of %v tickets and %v are still available", conferenceTickets, availableTickets)
	fmt.Println("Get your ticket here to attend")

	var firstName string
	var lastName string
	var email string
	var tickets int

	fmt.Println("Enter Your First Name \n")
	fmt.Scan(&firstName)
	fmt.Println("Enter Your Last Name \n")
	fmt.Scan(&lastName)
	fmt.Println("Enter Your Email Name \n")
	fmt.Scan(&email)
	fmt.Println("Enter Your Tickets_no \n")
	fmt.Scan(&tickets)
	fmt.Printf("Thank you %v   %v for booking %v tickets. You will receive email confirmation at %v \n", firstName, lastName, email, tickets)
	fmt.Printf("conferenceName type is %T and conferenceTickets type is %T", conferenceName, conferenceTickets)

}
