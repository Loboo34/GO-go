package main

import "fmt"

type People struct {
	id        int
	firstName string
	lastName  string
	age       int
	voted     bool
}

var voters []People
var nextId = 1

func main() {

	for {
		addUser()
	}

}

func userInput() (string, string, int) {
	var firstName string
	var lastName string
	var age int
	fmt.Println("Enter first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter Your age")
	fmt.Scan(&age)

	return firstName, lastName, age
}

func addUser() {
	firstName, lastName, age := userInput()

	var userData = People{
		id:        nextId,
		firstName: firstName,
		lastName:  lastName,
		age:       age,
		voted:     false,
	}

	voters = append(voters, userData)
	nextId++
	fmt.Printf("List of voters %v\n", voters)
}
