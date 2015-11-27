package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"time"
)

type Conf struct {
	test string `toml: "test"`
}

type Config struct {
	Age        int
	Cats       []string
	Pi         float64
	Perfection []int
	DOB        time.Time // requires `import time`
}

func main() {
	fd, err := os.Open("./example.toml")
	buf := make([]byte, 1024)
	fd.Read(buf)
	var conf Config

	_, err = toml.DecodeFile(buf, &conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(conf.test) < 0 {
		fmt.Println("error")
	}
	fmt.Print(conf)
}
