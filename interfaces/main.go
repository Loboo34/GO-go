package main

import (
	"fmt"

)

// type Laugh interface {
// 	Laugh() string
// }

// type Shape interface {
// 	Area() float64
// 	Perimeter() float64
// }

// type Circle struct {
// 	Radius float64
// }

// func (c Circle) Area() float64 {
// 	return 3.142 * c.Radius * c.Radius
// }

// // func (c Circle) Perimeter() float64 {
// // 	return 2 * 3.142 * c.Radius
// // }

// // type Cesar struct{}

// // func (c Cesar) Laugh() string {
// // 	return "Shirororo!shirororo!"
// // }

// func printShapeInfo(s Shape) {
//     fmt.Printf("Area: %.2f\n", s.Area())
//     //fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
// }

// func main() {

// 	fmt.Println("Learning about interfaces") //set of method signatures

// 	// var l Laugh = Cesar{}
// 	// fmt.Println(l.Laugh())

// 	c := Circle{Radius: 5}

// 	printShapeInfo(c)
// }


type Greeter interface{
	Greeting() string
}

type Person struct{
	Name string
	message string
}


func (p Person) Greeting() string{
	return fmt.Sprintf("hello %v, %v", p.Name,p.message)
}

func print(g Greeter){
	//fmt.Printf("Hello, %v",g.Greeting())
	fmt.Println(g.Greeting())
}


func main(){

	// gg := Person{Name: "Fred"}
	p1 := Person{
		Name: "Fred",
		message: "Have a great day",
	}
	print(p1)

}
