package hello_go

import "fmt"

func LenAndCap()  {
	mySlice := make([]string, 2, 5)
	mySlice[0] = "id-1"
	mySlice[1] = "id-2"

	fmt.Println("Demo of len and cap:")
	fmt.Println("len of init:", len(mySlice))
	fmt.Println("cap of init:", cap(mySlice))

	mySlice = append(mySlice, "id-3")
	mySlice = append(mySlice, "id-4")
	mySlice = append(mySlice, "id-5")
	mySlice = append(mySlice, "id-6")

	fmt.Println("len of append:", len(mySlice))
	fmt.Println("cap of append:", cap(mySlice))

	fmt.Println("")
}