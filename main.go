package main

import (
	"Hello-Go/hello_make"
	"fmt"
)

func main()  {
	helloGo()

	hello_make.MakeSlice()
	hello_make.MakeMap()
	hello_make.MakeChan()
}

func helloGo()  {
	fmt.Println("Hello, Go!")
	fmt.Println("")
}
