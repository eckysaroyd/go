package main

import (
	"strings"
)

func validate(firstName string, lastName string, email string, tickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := tickets > 0 && availableTickets >= tickets

	return isValidName, isValidEmail, isValidTicket
}
