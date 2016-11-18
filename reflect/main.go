// Author:         Wanghe
// Email:          wanghe@renrenche.com
// Author website: http://example.cn

// File: main.go
// Create Date: 2016-07-17 16:49:40

package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
}

type Bar struct {
}

var RegStruct map[string]interface{}

func main() {
	tr := "Bar"
	if RegStruct[tr] != nil {
		t := reflect.ValueOf(RegStruct[tr]).Type()
		v := reflect.New(t).Elem()
		fmt.Println(v)
	}
}

func init() {
	RegStruct = make(map[string]interface{})
	RegStruct["Foo"] = Foo{}
	RegStruct["Bar"] = Bar{}
}
