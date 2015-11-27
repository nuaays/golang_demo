package main

import (
	"fmt"
	"unsafe"
)

type T struct {
	A string
}

func main() {
	A string
	arr := make([]T, 0, 2)
	fmt.Println(len(arr))
	fmt.Println(unsafe.Sizeof(&A))
	arr = append(arr, T{"hello world"})

	fmt.Println(arr)
	fmt.Println(unsafe.Sizeof(arr))
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
}
