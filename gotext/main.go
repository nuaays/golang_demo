// Author:         Wanghe
// Email:          wanghe@renrenche.com
// Author website: http://example.cn

// File: main.go
// Create Date: 2016-08-07 09:04:21

package main
import (
					"github.com/leonelquinteros/gotext"
					"fmt"

)

func main() {
		gotext.Configure("/Users/wanghe/locales/", "zh_CN", "artemis")
		greet := gotext.Get("test")
		fmt.Println(greet)

}
