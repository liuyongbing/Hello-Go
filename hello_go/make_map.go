package hello_go

import "fmt"

func MakeMap() {
	m1 := map[string]string{
		"name": "init a map with values",
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
		"key4": "value4",
	}
	fmt.Println("Init a map with values:")
	fmt.Println(m1)
	fmt.Println("")

	// map key 是无序的
	fmt.Println("Traversing map: m1")
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	fmt.Println("")

	// m2 = nil
	var m2 map[int]string

	fmt.Println("var a map: the value is nil")
	fmt.Println(m2)
	fmt.Println("")

	// myMap == empty map
	myMap := make(map[int]string)

	myMap[10] = "Cat"
	myMap[20] = "Dog"
	myMap[30] = "Tiger"
	myMap[40] = "Fox"

	fmt.Println("Demo of make map: it is a empty map")
	fmt.Println(myMap)
	fmt.Println("")

	fmt.Println("Getting values")
	val, ok := myMap[10]
	fmt.Println(val, ok)
	val, ok = myMap[11]
	fmt.Println(val, ok)

	k := 100
	if val, ok := myMap[k]; ok {
		fmt.Println(val)
	} else {
		fmt.Printf("Key:%d does not exist", k)
		fmt.Println()
	}

	fmt.Println("Deleting values")
	fmt.Println(myMap)
	delete(myMap, 40)
	fmt.Println(myMap)
}

// func main() {
// 	MakeMap()
// }
