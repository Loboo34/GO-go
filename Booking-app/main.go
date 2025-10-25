package main

import (
	"fmt"
	//"strings"
	//"strconv"
	"time"
)

var plartformName = "AJX's"

// const tickets int = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	userEmail   string
	userTickets uint
}

func main() {

	greeting()

	for {
		//geting user input
		firstName, lastName, userEmail, userName, userTickets, err := getUserInput()
		if err != nil {
			fmt.Println("Unable to take user input")
			break
		}

		//input validation
		isValidName, isValidEmail, isValidTicket, err := inpuValidation(firstName, lastName, userEmail, userTickets)
		if err != nil {
			fmt.Println("Invalid input format", err)
			break
		}

		if isValidName && isValidEmail && isValidTicket {

			booking(userTickets, firstName, lastName, userEmail, userName)

			go printTicket(userName, userTickets, userEmail)

			printFirstName()

			if remainingTickets == 0 {
				fmt.Println("Tickets are sold out. Come back next time")
				break
			}
		}
		// } else {
		// 	//fmt.Println("Input Data is invalid. Please try again ")
		// 	if !isValidName {
		// 		fmt.Println("First name or last name is invalid")
		// 	}
		// 	if !isValidEmail {
		// 		fmt.Println("Wrong email format")
		// 	}
		// 	if !isValidTicket {
		// 		fmt.Printf("no of tickets %v is more than the available %v", userTickets, remainingTickets)
		// 	}
		// }
	}

}

// greeting function
func greeting() {
	fmt.Printf("Welcome to %v Booking.\n", plartformName)
	fmt.Println("Get Your Confrence Tickets here")
	fmt.Printf("%v available Tickets\n", remainingTickets)
}

// get user input func
func getUserInput() (string, string, string, string, uint, error) {
	var firstName string
	var lastName string
	var userEmail string
	var userTickets uint

	//taking user input
	fmt.Println("Enter your first name:")
	if _, err := fmt.Scan(&firstName); err != nil {
		return "", "", "", "", 0, err
	}

	fmt.Println("Enter your last name:")
	if _, err := fmt.Scan(&lastName); err != nil {
		return "", "", "", "", 0, err
	}

	var userName = firstName + " " + lastName

	fmt.Println("Enter your email address:")
	if _, err := fmt.Scan(&userEmail); err != nil {
		return "", "", "", "", 0, err
	}

	fmt.Println("How many tickets do you want to buy?")
	if _, err := fmt.Scan(&userTickets); err != nil {
		return "", "", "", "", 0, err
	}

	return firstName, lastName, userEmail, userName, userTickets, nil

}

// booking func
func booking(userTickets uint, firstName string, lastName string, userEmail string, userName string) UserData {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		userEmail:   userEmail,
		userTickets: userTickets,
	}
	// userData["firstName"] = firstName;
	// userData["lastName"] = lastName;
	// userData["userEmail"] = userEmail
	// userData["noOfTickets"]= strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of Bookings: %v\n", bookings)

	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, userTickets, userEmail)
	fmt.Printf("%v remaining tickets\n", remainingTickets)

	return userData
}

// printing first name func
func printFirstName() {
	firstNames := []string{}
	for _, vals := range bookings {

		//	firstName := names[0]

		firstNames = append(firstNames, vals.firstName)
	}
	fmt.Printf("Booking List: %v\n", firstNames)
}

func printTicket(userName string, userTickets uint, userEmail string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v ticket for %v", userTickets, userName)
	fmt.Println("#########")
	fmt.Printf("sending ticket:\n %v\n to %v\n", ticket, userEmail)
	fmt.Println("#########")

}
