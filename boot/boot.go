package main

import "fmt"

//Nested struct
type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

//embedded struct
type Address struct{
	city string
}

type Person struct{
	name string
	number int
	Address
}

type Auth struct{
	userName string
	password string
}





func main() {
	var msg messageToSend
	msg.sender.name = "fred"
	fmt.Println(msg.sender.name)

	var ps Person
	ps.Address = Address{city: "Mombasa"}
	fmt.Println(ps.Address)

	var auth Auth
	auth.userName = "fred"
	auth.password = "fre123"

	answ := ag(auth.userName, auth.password)

	fmt.Println(answ)


}

func ag(a string, b string) string{
return a + ":"+ b;

}
