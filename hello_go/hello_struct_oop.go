package hello_go

import "fmt"

type Animal struct {
	Color string
}

type Cat struct {
	Animal
	ID int
	Name string
	Age int
}

func (cat *Animal)Say()  {
	fmt.Println("Demo of struct oop:")
	fmt.Println("Miao miao")
	fmt.Println("")
}

func (cat *Cat) Run()  {
	fmt.Println("Cat ID:", cat.ID, "Cat is running.")
}
