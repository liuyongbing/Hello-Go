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

type Bird struct {
	Animal
	ID int
	Name string
	Age int
}

func (cat *Cat)Say() string {
	voice := "喵喵!"

	fmt.Println("Demo of struct oop:")
	fmt.Println(voice)
	fmt.Println("")

	return voice
}

func (cat *Cat)Run() string {
	starting := "Cat is running."
	fmt.Println("Cat ID:", cat.ID, starting)

	return starting
}

func (bird *Bird)Say() string {
	voice := "唧唧喳喳!"

	fmt.Println("Demo of interface:")
	fmt.Print(voice)
	fmt.Println("")

	return voice
}

func (bird *Bird)Run() string {
	starting := "Bird is running."

	fmt.Println("Bird ID:", bird.ID, starting)

	return starting
}
