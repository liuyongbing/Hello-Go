package hello_go

import "fmt"

// 指针
func GoPoint()  {
	var count = 20
	var countPoint *int
	countPoint = &count

	fmt.Println("Demo of point:")

	fmt.Println("count 变量的地址: ", &count)
	if countPoint != nil {
		fmt.Println("countPoint 变量存储的地址: ", countPoint)
		fmt.Println("countPoint 变量指针指向地址的值: ", *countPoint)
	}

	fmt.Println("")
}

// 指针数组
func PointArr()  {
	a,b := 1,2
	arr := [...]*int{&a, &b}
	fmt.Println("Dome of 指针数组:", arr)
	fmt.Println("")
}

// 数组指针
func ArrPoint()  {
	arr := [...]int{3, 4, 5}
	arrPoint := &arr

	fmt.Println("Dome of 数组指针:", arrPoint)
	fmt.Println("")
}
