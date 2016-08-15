// Author:         Wanghe
// Email:          wanghe@renrenche.com
// Author website: http://example.cn

// File: main.go
// Create Date: 2016-08-13 22:35:02

package main

import (
	"fmt"
)

func main() {
	total := 70
	per_page := 20
	slice := make([]int, total)
	for i := 0; i < len(slice); i++ {
		slice[i] = i
	}

	fmt.Println(slice[2:33])
	start := 0
	last_page_count := total % per_page
	if last_page_count == 0 {
		for start := 0; start < len(slice); start += per_page {
			fmt.Println(slice[start:per_page])
		}
		return
	}

	pages := total/per_page+1
	last_index := pages - 1
	for i := 0; i < total/per_page+1; i++ {
		start += per_page
		if start > len(slice) {
			break
		}
		fmt.Println(total / per_page)
		fmt.Println("i=", i)
		fmt.Println("pages=", pages)
		if i == 2 {
			fmt.Println("I am here")
			last_page := total % per_page
			fmt.Println(slice[start : start+last_page])
			break
		}

		fmt.Println(start)
		fmt.Println("start=", start, "per_page=", per_page)
		fmt.Println(slice[start : start+per_page])
	}

}
