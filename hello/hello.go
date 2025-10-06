package main

import "fmt"

var h = 4;

//%v-prints value of default format
//%T-prints the type of value
//%v#-Prints the value in Go-syntax format
//%%-Prints the % sign

func main(){
	//fmt.Println("sup bro!");

// 	var num int = 3;
// 	var student = "fred";
//   x := 4;
//   var b int;
//   var c string;
//   c= "yooh!"
//   b= 6;
// 	fmt.Println(num);
// 	fmt.Println(student);
// 	fmt.Println(x);
// 	fmt.Println(b);
// 	fmt.Println(c);
// 	fmt.Println(h);
// var a, b, c int = 2, 4, 6;

// fmt.Println(a)
// fmt.Println(b)
// fmt.Println(c)

// var d, e, f = 3, "free", true;

// fmt.Println(d)
// fmt.Println(e)
// fmt.Println(f)

// g, j := "take no", 30;
// fmt.Println(g)
// fmt.Println(j)

// var (
// 	a int = 34
// 	b int
// 	c = "fred")
// fmt.Println(a)
// fmt.Println(b)
// fmt.Println(c)
// const A = 3;

// const (
// 	B int = 6
// 	C = "Sam"
// )

// fmt.Println(A)
// fmt.Println(B)
// fmt.Println(C)

var a,b = "sup", true

//fmt.Print(a)//prints args in default format
//fmt.Print(b)
//fmt.Printf("%v is of type: %T \n", a, a)
//fmt.Printf("%v is of type: %T", b, b)
//fmt.Printf("%v is of type: %T\n", b, b)
fmt.Printf("%#v\n", b)
fmt.Printf("%v%%", a)
//fmt.Println(a,b)



}