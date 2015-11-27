package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := strings.NewReader("ABCDEFG")
	br := bufio.NewReader(s)
	c, _ := br.ReadByte()
	fmt.Printf("%c\n", c)
	// A

	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)
	// B

	f, _ := os.OpenFile("/opt/sda4/heka/wanghe-pc/access.log", os.O_RDWR, 0777)
	bw := bufio.NewWriter(f)

}
