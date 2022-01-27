package hello_go

import (
	"fmt"
	"reflect"
)

func NewMap() {
	myMap := new(map[int]string)

	fmt.Println("type of new map:", reflect.TypeOf(myMap))
}
