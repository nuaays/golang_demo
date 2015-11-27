package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 4, 5, 3, 1, 8, 9, 10}
	sort.Sort(sort.IntSlice(arr))
	fmt.Print(arr)
}
