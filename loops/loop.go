package main

import "fmt"

func main(){
// for   i := 0; i < 10;i++{
// 	fmt.Println("yooh")
// }

// for i := 0; i < 100; i+=10{
// if i == 30{
// 	continue;
// }
// fmt.Println(i)
// }

// for i := 5; i < 100; i+=10{
// 	if i == 55{
// 		break
// 	}
// 	fmt.Println(i)
// }

watchList := [3]string{"Kaiju no 8", "Silent witch", "One piece"};

for indx, value := range watchList {
	fmt.Printf("%v: %v\n", indx, value)
}

list := []string{"fred","mary","john"};

for _,value := range list{
	fmt.Printf("%v\n", value)
}
}