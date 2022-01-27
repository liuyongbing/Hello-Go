package hello_make

import "fmt"

func MakeSlice()  {
	mySlice := make([]string, 3)

	mySlice[0] = "Cat"
	mySlice[1] = "Dog"
	mySlice[2] = "Tiger"

	fmt.Print("Demo of make slice: ")
	fmt.Println(mySlice)
	fmt.Println("")
}