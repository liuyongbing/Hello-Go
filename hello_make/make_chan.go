package hello_make

import "fmt"

//创建没有缓存的 chan
func MakeChan()  {
	myChan := make(chan int)

	close(myChan)

	fmt.Print("Demo of make chan:")
	fmt.Println(myChan)
	fmt.Println("")
}
