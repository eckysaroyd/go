package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var availableTickets = 50
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("We have total of %v tickets and %v are still available", conferenceTickets, availableTickets)
	fmt.Println("Get your ticket here to attend")

	var userName string
	var tickets int

	userName = "Abbu"
	tickets = 50

	fmt.Printf(" %v has  %v tickets", userName, tickets)

}
