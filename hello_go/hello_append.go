package hello_go

import "fmt"

func SliceOfAppend()  {
	mySlice := make([]string, 2)
	mySlice[0] = "id-1"
	mySlice[1] = "id-2"

	fmt.Println("")
	fmt.Println("Demo of append:")

	// init of len(长度) & cap(容量)
	fmt.Println("len of init:", len(mySlice))
	fmt.Println("cap of init:", cap(mySlice))

	// append element
	mySlice = append(mySlice, "id-3")

	// append of len(长度) & cap(容量)
	fmt.Println("len of append:", len(mySlice))
	fmt.Println("cap of append:", cap(mySlice))

	fmt.Println("")
}