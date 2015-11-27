package main

import (
	"fmt"
	"strings"
)

func main() {
	filename := "/var/heka/StreamTable"
	arr := strings.Split(filename, "/")
	index := len(arr) - 1
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(arr[index])

}
