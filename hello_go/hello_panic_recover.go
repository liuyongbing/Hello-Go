package hello_go

import (
	"errors"
	"fmt"
)

// panic type: string
func PanicRecover()  {
	defer coverPanic()

	fmt.Println("Demo of panic and recover:")
	panic("I am panic(string)")
}

// panic type: error
func PanicRecover2()  {
	defer coverPanic()

	fmt.Println("Demo of panic and recover:")
	panic(errors.New("I am panic(object:error)"))
}

// panic type: unknown
func PanicRecover3()  {
	defer coverPanic()

	fmt.Println("Demo of panic and recover:")
	panic(1)
}

func coverPanic()  {
	message := recover()

	switch message.(type) {
		case string:
			fmt.Println("message type(string):", message)
		case error:
			fmt.Println("message type(error):", message)
		default:
			fmt.Println("message type(unknown):", message)
	}
	fmt.Println("")
}
