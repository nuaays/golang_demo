// Author:         Wanghe
// Email:          wanghe@renrenche.com
// Author website: http://example.cn

// File: main.go
// Create Date: 2016-07-17 19:26:47

package main

import "fmt"
import "os"

func main() {

	greet := os.Args[1]
	switch greet {
	case "hello":
		fmt.Println("hello")
		break
	case "world":
		fmt.Println("world")

	}
}
