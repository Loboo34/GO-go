package main

import "fmt"

func greeting(){
	fmt.Println("Yooh")
}

func sum(x int, y int){
	sm := x + y

	fmt.Println(sm)
}

func addTwoNumbers(a int, b int) (sum int){
	sum = a + b
	return
}

func main(){
greeting()
sum(2,3)
a := addTwoNumbers(4,2)
fmt.Println(a)
}