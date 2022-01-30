package hello_go

import "fmt"

type Dog struct {
	ID int
	Name string
	Age int
}

func (d *Dog)Run()  {
	fmt.Println("Dome of struct function()")
	fmt.Println("Obj of dog:", "ID:", "Dog is running.")
	fmt.Println("")
}

func StructDog()  {
	var dog Dog
	dog.ID = 1
	dog.Name = "旺财"
	dog.Age = 3

	fmt.Println("Demo of struct:")
	fmt.Println("Struct of dog(type:var):", dog)
	fmt.Println("")
}

func StructDog2()  {
	dog := Dog{ID: 2, Name: "来福", Age: 2}

	fmt.Println("Demo of struct:")
	fmt.Println("Struct of dog(type:':='):", dog)
	fmt.Println("")
}

func StructDog3()  {
	dog := new(Dog)
	dog.ID = 3
	dog.Name = "平安"
	dog.Age = 1

	fmt.Println("Demo of struct:")
	fmt.Println("Struct of dog(type:new()):", dog)
	fmt.Println("")
}

