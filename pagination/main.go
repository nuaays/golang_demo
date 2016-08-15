// Author:         Wanghe
// Email:          wanghe@renrenche.com
// Author website: http://example.cn

// File: main.go
// Create Date: 2016-08-15 11:08:47

package main

import (
	"flag"
	"fmt"
)

func main() {
	var data [50]int
	s := flag.Int("start", 0, "start")
	c := flag.Int("count", 0, "count")
	flag.Parse()
	for i := 0; i < 50; i++ {
		data[i] = i
	}

	start := *s
	count := *c

	end := start + count
	if end > len(data) {
		end = len(data)
	}
	fmt.Println(data[start:end])
}
