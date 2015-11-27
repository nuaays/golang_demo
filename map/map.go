package main

import (
	"fmt"
	//	"os"
)

type FileEvent struct {
	A string
	B string
	c string
}

func main() {
	var m map[string]string = make(map[string]string)
	m["one"] = "1111"
	m["two"] = "2222"
	m["three"] = "3333"
	fmt.Println(m["one"])
	fmt.Println(m["one"])
	fmt.Println(m["three"])

	var m1 map[string]FileEvent = make(map[string]FileEvent)
	m1["one"] = FileEvent{"AA", "BB", "CC"}
	m1["two"] = FileEvent{"AAAA", "BBBB", "CCCC"}
	fmt.Println(m1["one"])

	if value, erx := m["thone"]; erx {
		fmt.Println(value)

	} else {
		fmt.Println(erx)
	}
	value2 := m["thone"]
	fmt.Println(value2)
	if value, erx := m["two"]; erx {
		fmt.Println(value)
	} else {
		fmt.Println(erx)
	}

}
