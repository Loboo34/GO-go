package main

import "fmt"

type anime struct{
	name string
	episodes int
}

func main(){
	var a1 anime
	var a2 anime

	a1.name = "One piece"
	a1.episodes = 116

	a2.name = "Kaiju no 8"
	a2.episodes = 24

	
	fmt.Println("Name:", a1.name)
	fmt.Println("Episodes:", a1.episodes)
	fmt.Println("Name:", a2.name)
	fmt.Println("Episodes:", a2.episodes)
	
}
