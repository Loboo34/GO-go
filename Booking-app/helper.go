package main

import "strings"

//user input validation
func inpuValidation(firstName string, lastName string, userEmail string,  userTickets uint)(bool, bool, bool){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidTicket := userTickets > 0 && userTickets<=remainingTickets

	return isValidName, isValidEmail, isValidTicket
}