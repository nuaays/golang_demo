package main

import (
	"fmt"
)

type LV struct {
	length int
	value  string
}

func swap(a *interface{}, b *interface{}) {
	var t *interface{}
	*t = *a
	*b = *a
	*b = *t
}

func main() {
	arr := make([]interface{}, 0, 2)
	arr = append(arr, LV{2, "he"})
	arr = append(arr, 9)
	arr = append(arr, 9.2345)
	arr = append(arr, "string")

	switch v := arr[0].(type) {
	case LV:
		fmt.Println(v.length)
	case int:
		fmt.Println(v)
	}

	a := 1
	b := 2
	swap(a, b)
	fmt.Println(&a, &b)

}
