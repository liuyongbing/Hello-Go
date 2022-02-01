package hello_go

import (
	"fmt"
	"time"
)

func Loop()  {
	for i := 1; i < 11; i++ {
		time.Sleep(time.Microsecond * 10)
		fmt.Print(i, ",")
	}
}
