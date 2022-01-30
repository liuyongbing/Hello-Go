package hello_go

import "fmt"

func ChanClose()  {
	fmt.Println("Demo of close:")

	myChan := make(chan int, 1)
	defer close(myChan)

	myChan <- 1

	fmt.Println(myChan)
	fmt.Println("")
}
