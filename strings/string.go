package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	filename := "/var/heka/StreamTable"
	arr := strings.Split(filename, "/")
	index := len(arr) - 1
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(arr[index])
	array := make([]string, 0)
	for i := 0; i < 10; i++ {
		array = append(array, strconv.Itoa(i))
	}
	fmt.Println("array=", array)
	hello := strings.Join(array, ",")
	fmt.Println(hello)
}
