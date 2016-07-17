package main

import (
	"container/list"
	"fmt"
	"os"
	"path/filepath"

	"io/ioutil"
)

func GetFullPath(path string) string {
	absolutePath, _ := filepath.Abs(path)
	return absolutePath
}

func PrintFilesName(path string) {
	fullPath := GetFullPath(path)
	//listStr := list.New()
	filepath.Walk(fullPath, func(path string, fi os.FileInfo, err error) error {
		fmt.Println(fi.Name())
		return nil
	})

}


func GetFileNames(path string) *list.List{
  l := list.New()
  filepath.Walk(path, func(path string, file os.FileInfo, err error) error{
    fmt.Println("filename: ", file.Name() )
    l.PushBack(file)
		return nil
  })
  return l
}


func GetFileList(root string) *list.List {
	filelist , err := ioutil.ReadDir(root);
	if err != nil {
		return nil
	}
	for _, info := range filelist {
		fmt.Println(GetFullPath(info.Name()))
	}
	return nil
}


func main() {
	PrintFilesName("/Users/wanghe/work/dev/renrenche/docs")
	// l := GetFileNames("/Users/wanghe/work/dev/renrenche/docs")
	// for p := l.Front(); p != nil; p=p.Next() {
	// 	filename := p.Value.(os.FileInfo).Name()
	// 	fmt.Println(GetFullPath(filename))
	// }

	GetFileList("/Users/wanghe/work/dev/renrenche/docs")
}
