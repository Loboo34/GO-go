package main

import "fmt"

func main(){

	// x := 4

	// p := &x//stores mem-address of x

	// *p = 10//modifies value at the address p

	// fmt.Printf("x: %v\n",x)
	// fmt.Printf("p: %v\n", p)//mem address
	// fmt.Printf("%v", *p)//gives value stored in the address p


	// val := 15
	// point(&val)
	// fmt.Println(val)

	// var p *int
	// fmt.Println(p)//nil pointer
	x := 25
	update(&x)

	fmt.Println(x)

}

// func point(num *int){
// 	*num++
// }

func update(val *int){//parameter of name val, *int-pointer to an int
	*val = 30 //modifies the value at address *val points to
}