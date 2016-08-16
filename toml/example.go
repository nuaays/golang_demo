package main

import (
	"fmt"
	"github.com/pelletier/go-toml"
)

func main() {
	conf, err := toml.LoadFile("./example.toml")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	a := conf.Get("test.data").([]interface{})
	fmt.Println(a)

}
