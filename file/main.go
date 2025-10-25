package main

import (
	"fmt"
	"os"
)

// func createWrite(path string) *os.File{
// if _,err :=os.Stat(path);os.IsNotExist(err){//os.Stat(path)-checks if file exists, if file does not exist create it
// 	f,err := os.Create(path)//creates new file
// 	if err != nil{
// 		panic(err)
// 	}
// 	return f
// }
// //open file
// f, err := os.OpenFile(path, os.O_RDWR, 0644)//open existing file for read&write. o_RDWR=read and write mode, 0644-permissions:owner can R&W while others just R
// if err != nil {
// 	panic(err)
// }
// return f
// }

func createWrite(path string) *os.File{
	f,err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return f
}

func main(){
	fp := createWrite("test.txt")//call function to create/open file whith name "test.txt"
	// fp.WriteString("how are you")//write text in file
	// fp.Close()//closes the file

	defer fp.Close()

	//write file
	_,err := fp.WriteString("Good\n")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Wrote to file successfully!")

	wr := create("text.txt")

	str := "Yeah!!"

	for i := 0; i < 10;i++{
		//fmt.Println(str)
		_,err = wr.WriteString(str)
	if err != nil {
		fmt.Print("Error",err)
	}
	}

	




}

func create(path string)*os.File{
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil{
		panic(err)
		
	}
	return f

}