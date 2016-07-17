// 使用[]rune 来处理utf8, 可以用[]rune 下标来访问单个中文字符
// 汉字在utf8中占３个字节
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	//	"strconv"
	"unicode/utf8"
)

func main() {
	//中文处理
	str := "百泉众合"
	fmt.Println(len(str))
	r := []rune(str)
	fmt.Println(r)
	for i := 0; i < len(r); i++ {
		fmt.Println(string(r[i]))
	}

	fmt.Println("-----------split line-----------")

	// utf8
	var string_value string = "hello, 世界"
	//	var int_value int = 1
	var float_value float32 = 2.0
	//	var double_value float64 = 3.3453
	b := []byte(string_value)
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		fmt.Printf("%c, %v\n", r, size)
		b = b[size:]
	}

	// string
	fmt.Println("-----------split line-----------")
	for len(string_value) > 0 {
		r, size := utf8.DecodeRuneInString(string_value)
		fmt.Printf("%c, %v\n", r, size)
		string_value = string_value[size:]
	}
	// float32
	fmt.Println("-----------split line-----------")
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, float_value)
	b = buf.Bytes()

	fmt.Print(buf.Bytes())


	fmt.Println("hell ----------------------------------------")
	fmt.Println("hell ----------------------------------------")
	fmt.Println("hell ----------------------------------------")


}
