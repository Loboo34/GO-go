package main

import "fmt"

func main(){
                               //arrays//
//var myNums = [7]int{1,2:2,3,4};
//var nums = [...]int{4,5,6,7,8,9}

//studs := [...]string{"fred", "mary","john"};
//watchList :=[4]string{"Onepiece", "Bleach", "slime"}

//watchList[1] = "JJK";


	//fmt.Println(myNums)
	//fmt.Println(nums)
	//fmt.Println(studs)
	//fmt.Println(len(watchList))
	//fmt.Println(watchList[0])


    //                            //Slices//
	// watchList := []string{"JJK", "slime","my dress up darling", "Gachiakuta", "Kaiju no 8"}

	// fmt.Println(watchList[0])
	// fmt.Println(len(watchList))
	// fmt.Println(cap(watchList))

	// pick := watchList[2:5]
	// pick = append(pick, "one piece", "sakamoto days")

	// fmt.Println(pick)

	// list := make([]string, 4, 100);

	// list = append(list, "fred", "james")

	// fmt.Println(list)
	// fmt.Println(cap(list))
	// fmt.Println(len(list))

	const (
		a = iota +2
		b =  iota
		c
		d
	)

	fmt.Println(a, b, c, d)
}