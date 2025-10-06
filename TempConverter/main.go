package main

import ("fmt"
"strings"
)

	var temp float32
	
	var choice string

	//to farh convert = (temp * 9) /5 + 32;
	//to celcius  convert = (temp - 32)* 5/9;
func main(){
	fmt.Println("Welcome to your go temp converter °")

	getUserInput()

	fmt.Println("Enter the temperature")
	fmt.Scan(&temp)


	switch choice {
		case "celsius":
			result :=convertToCelsius()
			fmt.Printf("%.2f °f converted to celsius is: %.2f °c\n", temp, result)
		case "fahrenheit":
			result := convertToFahrenheit()
			fmt.Printf("%.2f °c converted to fahrenheit is: %.2f °f\n", temp, result)

		default:
			fmt.Println("Invalid choice")
	}


}

func getUserInput(){
	fmt.Println("what would you like to convert to? celsius or fahrenheit")
	fmt.Scan(&choice)
	choice = strings.ToLower(strings.TrimSpace(choice))
}

func convertToCelsius() float32{
	return (temp - 32) * 5/9
}
func convertToFahrenheit() float32{
	return (temp * 9)/5 +32
}