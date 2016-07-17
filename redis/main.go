// Author:         Wanghe
// Email:          wanghe@renrenche.com
// Author website: http://example.cn

// File: main.go
// Create Date: 2016-07-13 11:55:34

package main

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
	"unicode/utf8"
	"gopkg.in/redis.v3"
)

func GzipEncode(in []byte) ([]byte, error) {
	var (
		buffer bytes.Buffer
		out    []byte
		err    error
	)
	writer := gzip.NewWriter(&buffer)
	_, err = writer.Write(in)
	if err != nil {
		writer.Close()
		return out, err

	}
	err = writer.Close()
	if err != nil {
		return out, err

	}

	return buffer.Bytes(), nil

}

func GzipDecode(in []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(in))
	if err != nil {
		var out []byte
		return out, err

	}
	defer reader.Close()

	return ioutil.ReadAll(reader)

}

//进行zlib压缩
func DoZlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(src)
	w.Close()
	return in.Bytes()

}

//进行zlib解压缩
func DoZlibUnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}

func Mget(client *redis.Client, carid int) {
	key := fmt.Sprintf("car:detail:%d", carid)
	bytes, _ := client.Get(key).Bytes()
	//	var f interface{}
	b := DoZlibUnCompress(bytes)
	fmt.Println(string(b))
	var r rune
	var n int
	r, n = utf8.DecodeRune(b)
	fmt.Printf("r=%s, %d\n", r, n)

}


func TestArray(a *[]int) {
				for _, e := range *a {
								fmt.Println(e)
				}
}


func main() {
	//address := "e3b8d1e0b3c940a3.m.cnbja.kvstore.aliyuncs.com:6379"
	//password := "e3b8d1e0b3c940a3:5UlzBT8bg0LNqe3ZrXmPVpEi"
	address := "10.46.176.68:6379"
	password := "P8NZcVSa"
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	pong, _ := client.Ping().Result()
	fmt.Println("pong=", pong)

	// 批量获取车辆详情数据
	Mget(client, 661329)

	arr := make([]int, 0)
	for i := 0; i<10; i++ {
					arr = append(arr, i)
	}

	for _, a := range arr {
		fmt.Println(a)	
	}

	fmt.Println("arr----------------")
	TestArray(&arr)
	//b, _ := GzipEncode([]byte("hello world"))
	//c, _ := GzipDecode(b)
	//fmt.Println(string(c))
}
