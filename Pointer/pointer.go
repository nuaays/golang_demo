package main

import (
	"fmt"
)

func main() {
	buff1 := make([]byte, 15)
	buff2 := make([]byte, 15)
	copy(buff1, []byte("hello world"))
	fmt.Println("buff1:= ", buff1)
	buff2 = buff1
	fmt.Println("buff2:=", buff2)

	buff3 := new([]byte)
	buff4 := new([]byte)
	buff3 = &buff1
	buff4 = buff3
	fmt.Println("buff3=", buff3)
	fmt.Println("buff4=", buff4)

	buff5 := new([15]byte)
	fmt.Println("buff5=", buff5)

}
