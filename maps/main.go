package main

import "fmt"

// func main() {
// 	//map[keytype]valuetype
func drink() {
	drinks := make(map[string]string)
	drinks["Soda"] = "coke"
	drinks["tea"] = "green"
	drinks["juice"] = "mango"

	drinks["Soda"] = "fanta"
	delete(drinks, "tea")

	for name, val := range drinks{
		fmt.Printf("%v : %v", name, val)
	}
	
}

// 	ages := map[string]int{
// 		"Luffy": 19,
// 		"Zoro":  21,
// 		"Nami":  20,
// 	}

// 	positions := make(map[string]string)
// 	positions["Captain"] = "Luffy"
// 	positions["Navigator"] = "Nami"

// 	ages["Luffy"] = 20//update value
// 	positions["Doctor"] = "chopper"
// 	positions["Comedian"] = "Bonclay"
// 	positions["Right Hand"] = "Zoro"

// 	//checking if key exist
// 	position, exists := positions["Right Hand"]
// 	if exists {
// 		fmt.Printf("Rignt hand man is %v\n", position)
// 	} else {
// 		fmt.Printf("Right hand man not yet joined the crew\n")
// 	}

// 	//delete(positions, "Comedian")
// 	// fmt.Println(ages["Luffy"])
// 	// fmt.Println(positions["Captain"])
// 	// fmt.Println(positions["Comedian"])

// 	for position, name := range positions {
// 		fmt.Printf("The %v is: %v\n", position, name) //the order is random
// 	}

// }

func main() {
	count := make(map[string]int)
	names := []string{"Luffy", "zoro", "Nami", "Ussop", "Sangi", "Choper", "Robin", "Brook", "zoro", "Sanji", "zoro"}

	for _, word := range names {
		count[word]++
	}

	fmt.Println(count)

	crew := map[string]map[string]string{
		"Luffy": {
			"role": "captain",
			"age":  "19",
		},
		"Zoro": {
			"role": "SwordsMan",
			"age":  "21",
		},
		"Nami": {
			"role": "Navigator",
			"age":  "20",
		},
	}

	for name, role := range crew {
		fmt.Printf("Name:%v| Role: %v \n", name, role)
	}

	drink()

}
