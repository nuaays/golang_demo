package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Conf struct {
	test string `toml: "test"`
}

func main() {
	var conf Conf
	_, err := toml.DecodeFile("/home/wanghe/work/toml/example.toml", &conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(conf.test) < 0 {
		fmt.Println("error")
	}
}
