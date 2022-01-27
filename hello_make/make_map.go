package hello_make

import "fmt"

func MakeMap()  {
	myMap := make(map[int]string)

	myMap[10] = "Cat"
	myMap[20] = "Dog"
	myMap[30] = "Tiger"

	fmt.Print("Demo of make map:")
	fmt.Println(myMap)
	fmt.Println("")
}