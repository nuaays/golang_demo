package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.MkdirAll("/var/heka", 0755)
	if err != nil {
		fmt.Println("Create dir error:", "/var/heka")
	}

	var fd *os.File
	fd, err = os.OpenFile("/var/heka/StreamTable", os.O_RDWR, 0660)
	buff := make([]byte, 2048)
	_, err = fd.Read(buff)
	fmt.Println(buff)
}
