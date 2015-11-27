package main

import (
	"flag"
	//	"fmt"
	"os"
	"time"
)

func write_buff(fd *os.File, buff []byte) {
	t := time.NewTicker(1 * time.Second)
	for {
		if fd != nil && len(buff) != 0 {
			fd.Write(buff)
		}
		select {
		case <-t.C:
			continue
		}
	}
}

func main() {
	filename := flag.String("s", "./log", "log file")
	flag.Parse()
	filename1 := "./1.log"
	filename2 := "./2.log"
	filename3 := "./3.log"
	filename4 := "./4.log"

	var (
		fd  *os.File
		fds [4]*os.File
		err error
	)

	fd, err = os.OpenFile(*filename, os.O_RDWR, 0660)
	fds[0], err = os.OpenFile(filename1, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	fds[1], err = os.OpenFile(filename2, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	fds[2], err = os.OpenFile(filename3, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	fds[3], err = os.OpenFile(filename4, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)

	buff := make([]byte, 2048)
	for i := 0; i < 4; i++ {
		_, err = fd.Read(buff)
		if err != nil {
			return
		}
		go func() {
			for i := 0; i < 4; i++ {
				go write_buff(fds[i], buff)
			}
		}()
	}
}
