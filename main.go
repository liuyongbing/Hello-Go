package main

import (
	"Hello-Go/hello_go"
	"fmt"
)

func main() {
	// Hello Go
	helloGo()

	// 内建方法
	helloMake()

	// struct
	helloStruct()

	// interface
	helloInterface()
}

// Interface
func helloInterface()  {
	fmt.Println("Demo of interface:")

	cat := new(hello_go.Cat)
	bird := new(hello_go.Bird)
	action(cat)
	action(bird)

	fmt.Println("")
}

// Interface
func action(b hello_go.Behavior) string {
	b.Run()
	b.Say()

	return ""
}

// Struct
func helloStruct()  {
	// struct
	hello_go.StructDog()
	hello_go.StructDog2()
	hello_go.StructDog3()

	// struct function
	dog := new(hello_go.Dog)
	dog.ID = 1000
	dog.Run()

	// struct oop
	cat := new(hello_go.Cat)
	cat.ID = 2000
	cat.Color = "White"
	cat.Run()
	cat.Say()
}

// 内建方法
func helloMake()  {
	// 内建方法: make()
	hello_go.MakeSlice()
	hello_go.MakeMap()
	hello_go.MakeChan()

	// 内建方法: new()
	hello_go.NewMap()
	hello_go.MakeMap2()

	// 内建方法: append()
	hello_go.SliceOfAppend()

	// 内建方法: copy()
	hello_go.SliceOfCopy()

	// 内建方法: delete()
	hello_go.MapOfDelete()

	// 内建方法: panic() & recover()
	hello_go.PanicRecover()
	hello_go.PanicRecover2()
	hello_go.PanicRecover3()

	// 内建方法: len() & cap()
	hello_go.LenAndCap()

	// 内建方法: close()
	hello_go.ChanClose()
}

// Hello Go
func helloGo() {
	fmt.Println("Hello, Go!(2022-01-27)")
	fmt.Println("")
}
