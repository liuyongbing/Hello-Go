package hello_go

import "fmt"

func MapOfDelete()  {
	myMap := make(map[int]string)
	myMap[0] = "id-1"
	myMap[1] = "id-2"

	fmt.Println("Demo of delete:")
	fmt.Println("删除之前: ", myMap)

	delete(myMap, 1)

	fmt.Println("删除之后: ", myMap)

	delete(myMap, 0)

	fmt.Println("再删之后: ", myMap)
	fmt.Println("")
}