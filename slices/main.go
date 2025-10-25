package main

import "fmt"

// func main(){
// 	//myArray := [2]string{"Luffy","Zoro"}//array with 2 elements
// 	//myArray1 := [...]int{1, 2,2,4}//array with unUndefined number of elements

// 	// a := myArray1
// 	// a[0] = 33
// 	//myArray1[4] = 44invalid

// 	//fmt.Println(myArray)
// 	// fmt.Println(a)
// 	// fmt.Println(myArray1)

//slices
// 	//myNums := []int{2,5,6,8,9}
// 	//myList := []string{"Luffy", "zoro","Nami"}
// 	myWatchlist := make([]string, 0)
// 	myTop := make([]string, 0)

// 	//myNums = append(myNums, 10, 13)
// 	//myList = append(myList, "Ussop", "Sanji")
// 	myWatchlist = append(myWatchlist, "Onepiece", "to your eternity")
// 	myTop = append(myTop, "Overlord", "Tensura", "Horimiya")

// 	myList := append(myTop, myWatchlist...)

// 	streaming := append(myTop, myWatchlist...)

// 	streaming = append(streaming[:2], streaming[3:]...)

// 	// fmt.Println(myNums)
// 	 fmt.Println(myList)
// 	// fmt.Println(myWatchlist)
// 	// fmt.Println(myTop)
// 	fmt.Println(streaming)
// }

func main() {
	nums := [3]int{1, 3, 5}
	mod(&nums)

	fmt.Println(nums)
}

func mod(arr *[3]int) {
	arr[0] = 0
}
