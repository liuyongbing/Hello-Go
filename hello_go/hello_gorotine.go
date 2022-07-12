// package hello_go
package hello_go

import (
	"fmt"
	"runtime"
	"time"
)

func Loop() {
	for i := 1; i < 11; i++ {
		time.Sleep(time.Microsecond * 10)
		fmt.Print(i, ",")
	}
}

func Loop2(i int) {
	fmt.Printf("Hello from goroutine %d\n", i)
}

func main() {
	var a [10]int
	count := 10
	for i := 0; i < count; i++ {
		go func(i int) {
			for {
				a[i]++
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
