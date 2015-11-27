package main

import (
	"fmt"
	//	"time"
)

var ch chan string = make(chan string, 10)

func listen(ch1 chan string) {
	for {
		select {
		case buf := <-ch1:
			fmt.Println("debug")
			fmt.Println(buf)
			/*
				case <-time.After(900000):
					fmt.Println("time out")
					return
			*/
		default:
		}
	}
}

func main() {
	greet := "hello go"
	ch <- greet
	go listen(ch)
	for {
		fmt.Scan(&greet)
		ch <- greet
	}
	// str := <-ch
	//fmt.Println(str)

}
