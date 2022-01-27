package hello_make

import (
	"fmt"
	"reflect"
)

func MakeMap2()  {
	myMap := make(map[int]string)

	fmt.Println("type of make map:", reflect.TypeOf(myMap))
}

