package main

import (
	"fmt"
)

var ch chan []int
var m []int

func main() {
	ch = make(chan []int)
	go func() {
		for {
			select {
			case c := <-ch:
				fmt.Println(c)
				fmt.Println("hello world")
			}
		}
	}()

	var i int
	fmt.Scanf("%d", &i)
	m = append(m, i)
	//	ch <- m

}
