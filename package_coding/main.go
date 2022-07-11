package main

import (
	"fmt"
	"sort"

	"github.com/liuyongbing/Hello-Go/strsort"
)

func main() {
	ss := []string{"liu", "fei", "lia"}
	// msort := strsort.NewMSort(ss)
	msort := strsort.NewMSort(ss)
	sort.Sort(msort)
	fmt.Println(msort)
}
