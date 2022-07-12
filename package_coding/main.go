package main

import (
	"fmt"
	"sort"

	"github.com/liuyongbing/Hello-Go/package_coding/multiple_strings_sort"
)

func main() {
	ss := []string{"liu", "fei", "lia"}
	// msort := strsort.NewMSort(ss)
	msort := multiple_strings_sort.NewMSort(ss)
	sort.Sort(msort)
	fmt.Println(msort)
}
