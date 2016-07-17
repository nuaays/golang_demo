// Author:         Wanghe
// Email:          wanghe@renrenche.com
// Author website: http://example.cn

// File: main.go
// Create Date: 2016-07-04 19:59:51

package main

import (
	"fmt"
	"time"
)

// 格式化当前时间
func GetCurrentDatetime() string {
	return time.Now().Format("2006-01-02 03:04:05")
}

// 时间戳转化为格式化时间
func GetDatetimeFormat(ts int64 ) string {
	return time.Unix(1389058332, 0).Format("2006-01-02 15:04:05")
}


// 格式化的时间转化为时间戳
func GetTimestampUnix(ts string) int64 {
				t := time.Parse("2006-01-02 03:04:05", "2014-01-03 33:23:53")
				return t.Unix()
}
func main() {
	fmt.Println(GetCurrentDatetime())
	fmt.Println(GetDatetimeFormat(1389058332))
	fmt.Println(GetTimestampUnix())
}


