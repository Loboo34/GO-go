package main

import (
	"fmt"
	"strings"
)

//user input validation
func inpuValidation(firstName string, lastName string, userEmail string,  userTickets uint)(bool, bool, bool, error){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2;
	 if !isValidName{
		return false, false, false, fmt.Errorf("inValid name %v %v",firstName, lastName)
	}
	isValidEmail := strings.Contains(userEmail, "@")
	 if !isValidEmail{
		return true, false, false, fmt.Errorf("inValid email %v ",userEmail)
	}
	isValidTicket := userTickets > 0 && userTickets<=remainingTickets
	 if !isValidTicket{
		return false, true, false, fmt.Errorf("invalid ticket count %v",userTickets)
	}

	return isValidName, isValidEmail, isValidTicket, nil
}