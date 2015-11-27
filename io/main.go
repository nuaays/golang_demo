package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type A struct {
	Offset int `json:"offset"`
}

func main() {
	path := "/home/wanghe/workspace/work/VerificationCode/io/a.txt"
	fd, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0660)
	if err != nil {
		fmt.Println(err)
	}
	defer fd.Close()
	var i int
	var a A
	for {
		i++
		a.Offset = i
		b, _ := json.Marshal(a)
		fd.Write(b)
		time.Sleep(time.Second)
	}
}
