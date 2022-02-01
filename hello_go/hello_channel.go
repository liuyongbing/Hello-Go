package hello_go

import (
	"fmt"
	"sync"
	"time"
)

var myChan = make(chan int, 10)
var timeOut = make(chan bool)

var WG sync.WaitGroup

// 协程: 发送数据
func Send()  {
	time.Sleep(time.Second * 1)
	fmt.Println("协程发送数据 num:" , 1)
	myChan <- 1

	time.Sleep(time.Second * 1)
	fmt.Println("协程发送数据 num:" , 2)
	myChan <- 2

	time.Sleep(time.Second * 1)
	fmt.Println("协程发送数据 num:" , 3)
	myChan <- 3

	time.Sleep(time.Second * 1)
	fmt.Println("协程发送数据 data:" , true)
	timeOut <- true
}

// 协程: 接收数据
func Receive()  {
	//num := <- myChan
	//fmt.Println("协程接收数据 num:" , num)
	//num = <- myChan
	//fmt.Println("协程接收数据 num:" , num)
	//num = <- myChan
	//fmt.Println("协程接收数据 num:" , num)
	for  {
		select {
			case num := <- myChan:
				fmt.Println("协程接收数据 num:", num)
			case tData := <- timeOut:
				fmt.Println("协程接收数据 data:", tData)
		}
	}
}


func Read()  {
	for i:=0; i<3; i++ {
		WG.Add(1)
	}
}

func Write()  {
	for i:=0; i<3; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println("Writed:", i)
		WG.Done()
	}
}

