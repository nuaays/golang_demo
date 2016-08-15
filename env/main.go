// Author:         Wanghe
// Email:          wanghe@renrenche.com
// Author website: http://example.cn

// File: main.go
// Create Date: 2016-07-17 17:09:20


package main

import (
	"fmt"
	"os"
)

func main() {
				debug := os.Getenv("debug1")
				fmt.Println(debug)
}



