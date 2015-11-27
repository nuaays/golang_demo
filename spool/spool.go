package main

import (
	"fmt"
)

var buffer chan string

func main() {
	buffer = make(chan string, 10)
	fmt.Println("start")
	go func() {
		fmt.Println("write")
		for {
			b := "abcdef"
			buffer <- b //"abcde"
			//buffer = append(buffer, b)
			//fmt.Println(<-buffer)
		}
	}()
	go func() {
		fmt.Println("read")
		for {
			select {
			case s := <-buffer:
				fmt.Println(s)
			}
		}
	}()
	for {
	}
}
