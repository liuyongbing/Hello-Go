package hello_go

import "fmt"

func SliceOfCopy()  {
	sliceSrc := make([]string, 2)
	sliceSrc[0] = "Slice src:1"
	sliceSrc[1] = "Slice src:2"

	sliceDst := make([]string, 3)
	sliceDst[0] = "Slice dst:1"
	sliceDst[1] = "Slice dst:2"
	sliceDst[2] = "Slice dst:3"

	fmt.Println("")
	fmt.Println("Demo of copy:")
	fmt.Println("copy 之前的 sliceSrc:", sliceSrc)
	fmt.Println("copy 之前的 sliceDst:", sliceDst)

	copy(sliceDst, sliceSrc)

	fmt.Println("copy 之后的 sliceDst:", sliceDst)

	fmt.Println("")
}