package main

import (
	"Hello-Go/hello_go"
	"fmt"
)

func main() {
	helloGo()

	// 内建方法: make()
	hello_go.MakeSlice()
	hello_go.MakeMap()
	hello_go.MakeChan()

	// 内建方法: new()
	hello_go.NewMap()
	hello_go.MakeMap2()
}

func helloGo() {
	fmt.Println("Hello, Go!(2022-01-27)")
	fmt.Println("")
}
