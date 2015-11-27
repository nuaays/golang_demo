package main

import (
	"fmt"
	"os"
	"regexp"
)

func MkDatadir(defaultbasedir string, user string, segment_count int, logfile_count int) []string {
	root := "data"
	paths := make([]string, 0)
	userdir := fmt.Sprintf("%s/%s/%s", defaultbasedir, user, root)
	for j := 0; j < segment_count; j++ {
		segmentdir := fmt.Sprintf("%s/%s%d", userdir, "segment", j)
		if err := os.MkdirAll(segmentdir, 0777); err != nil {
			fmt.Println(err)
		}
		for k := 0; k < logfile_count; k++ {
			logfile := fmt.Sprintf("%s/%d", segmentdir, k)

			paths = append(paths, logfile)
			//		fd, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
			//		defer fd.Close()
			//		if err != nil {
			//			fmt.Println(err)
			//		}
		}
	}
	return paths

}

func MkIndexdir(defaultbasedir string, user string, segment_count int, logfile_count int) []string {
	root := "indexes"
	paths := make([]string, 0)
	userdir := fmt.Sprintf("%s/%s/%s", defaultbasedir, user, root)
	for j := 0; j < segment_count; j++ {
		segmentdir := fmt.Sprintf("%s/%s%d", userdir, "segment", j)
		if err := os.MkdirAll(segmentdir, 0777); err != nil {
			fmt.Println(err)
		}
		for k := 0; k < logfile_count; k++ {
			logfile := fmt.Sprintf("%s/%d", segmentdir, k)
			paths = append(paths, logfile)
			//	fd, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
			//	defer fd.Close()
			//	if err != nil {
			//		fmt.Println(err)
			//	}
		}
	}
	return paths

}

func main() {
	defaultbasedir := "/home/wanghe/workspace/work/VerificationCode/MakeVirtualFile"
	user := "user0"
	segment_count := 2
	logfile_count := 3
	data_paths := MkDatadir(defaultbasedir, user, segment_count, logfile_count)
	index_paths := MkIndexdir(defaultbasedir, user, segment_count, logfile_count)
	fmt.Println(data_paths, index_paths)
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

}
