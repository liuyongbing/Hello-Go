package main

import (
	"Hello-Go/hello_make"
	"fmt"
)

func main()  {
	helloGo()

	// 内建方法: make()
	hello_make.MakeSlice()
	hello_make.MakeMap()
	hello_make.MakeChan()

	// 内建方法: new()
	hello_make.NewMap()
	hello_make.MakeMap2()
}

func helloGo()  {
	fmt.Println("Hello, Go!(2022-01-27)")
	fmt.Println("")
}
