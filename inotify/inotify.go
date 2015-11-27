package main

import (
	"flag"
	"fmt"
	"golang.org/x/exp/inotify"
	"log"
	"os"
	"strings"
	"syscall"
)

func main() {
	logFileName := flag.String("log", "inotify.log", "Log file name")
	watchFileName := flag.String("watch", "/tmp", "Watch file name")
	flag.Parse()
	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "inotify server start Failed")
		os.Exit(1)
	}
	log.SetOutput(logFile)
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	temp := strings.Split(*watchFileName, "/")
	stream := strings.Split(temp[len(temp)-1], ".")[0]
	log.Println("stream_name: ", stream)

	streamFile, streamErr := os.OpenFile(stream, os.O_CREATE|os.O_RDWR, 0666)
	if streamErr != nil {
		fmt.Println("Fail to find", streamFile, "inotify server start FAiled")
		os.Exit(1)
	}
	// defer os.Close(streamFile)
	watcher, err := inotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	err = watcher.Watch(*watchFileName)
	if err != nil {
		log.Fatal(err)
	}

	for {

		select {
		case ev := <-watcher.Event:

			switch ev.Mask {
			case syscall.IN_MOVE_SELF:
				log.Println("rename")
				break
			case syscall.IN_MOVED_TO:
				log.Println("move to")
				break
			case syscall.IN_MOVE:
				log.Println("move")
				break
			case syscall.IN_DELETE:
				log.Println("delete")
				break
				//		case syscall.IN_MODIFY:
				//			log.Println("modify")
				//			break
			case syscall.IN_CREATE:
				log.Println("Create")
				break
			case syscall.IFF_TUN_EXCL:
				log.Println("TRUNCAT")
				break
			}
			log.Println(ev)
			log.Println("mask:", ev.Mask, "name:", ev.Name, "Cookie:", ev.Cookie)
		case err := <-watcher.Error:
			log.Println("error:", err)
		}
	}
}
