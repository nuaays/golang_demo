/****************************************************************************
 *
 * 扫描指定的目录中存在的ＴＯＭＬ文件
 * 将这些ＴＯＭＬ文件合并生成一个ＭＤ５码
 * author: wanghe
 * email: wangh@loginsight.cn
 * date: 2015-8-12
 *
 ***************************************************************************/

package main

import (
	"crypto/md5"
//	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func GeneralMd5(str string) []byte {
	if len(str) <= 0 {
		return nil
	}
	h := md5.New()
	io.WriteString(h, str)
	return h.Sum(nil)
}

func IsContain(arr [6]string, e string) bool {
	var c int
	for _, elem := range arr {
		if elem != e {
			c++
			continue
		}
	}
	if c == len(arr) {
		return false
	}
	return true
}

func MergeTomlFile(tomlfiles []string) string {
	if len(tomlfiles) <= 0 {
		fmt.Errorf("tomlfiles does not exsits")
		return ""
	}
	var contents string
	for _, f := range tomlfiles {
		content, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Errorf("Failed to read file: '%s'", err.Error())
			return ""
		}
		contents += string(content)
	}
	return contents

}

func ScanTomlFile(config_path *string) []string {
	files := make([]string, 0, 100)
	if config_path == nil {
		return nil
	}
	dirs := ScanConfigDir(config_path)
	for _, d := range dirs {
		p, err := os.Open(d)
		if err != nil {
			fmt.Errorf("%s: faild to open file '%s' \n", d, err.Error())
			return nil
		}
		fi, err := p.Stat()
		if err != nil {
			fmt.Errorf("can not stat file: %s", err.Error())
			return nil
		}

		if fi.IsDir() {
			fs, _ := ioutil.ReadDir(d)
			for _, f := range fs {
				fName := f.Name()
				if strings.HasSuffix(fName, ".toml") {
					file_path := fmt.Sprintf("%s/%s", d, f.Name())
					files = append(files, file_path)
					fmt.Println(file_path)
				}
			}
		}
	}
	return files
}
func ScanConfigDir(config_path *string) []string {
	p, err := os.Open(*config_path)
	if err != nil {
		fmt.Printf("error opening file: %s", err.Error())
		return nil
	}

	fi, err := p.Stat()
	if err != nil {
		fmt.Printf("can't stat file: %s", err.Error())
		return nil
	}
	conf_list := [6]string{"active", "auto", "global", "available", "auto", "archive"}
	file_list := make([]string, 0, 100)
	if fi.IsDir() {
		elements, _ := ioutil.ReadDir(*config_path)
		for _, e := range elements {
			if IsContain(conf_list, e.Name()) {
				file_list = append(file_list, fmt.Sprintf("%s/%s", *config_path, e.Name()))
			}
		}
		return file_list
	}
	return nil

}
func main() {

//					config_path := flag.String("config", "/home/wanghe/work/heka/build/heka/conf", "config file or directory")
//	flag.Parse()
//	files := ScanTomlFile(config_path)
//	fmt.Println(files)
//	contents := MergeTomlFile(files)
	//fmt.Println(contents)
	contents := "ffffffffff"
	_md5 := GeneralMd5(contents)
	fmt.Printf("%x\n", _md5)
}
