package main

import (
	//	notify "github.com/rafrombrc/go-notify"
	"github.com/bbangert/toml"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	sigChan := make(chan os.Signal)
	log.Println("pid==", os.Getpid())
	pid := os.Getpid()
	process, err := os.FindProcess(pid)
	if err != nil {
		log.Println(err)
	}

	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT)

	for {
		select {
		case sig := <-sigChan:
			switch sig {
			case syscall.SIGHUP:
				log.Println("SIGNAL SIGHUP")
			case syscall.SIGINT:
				log.Println("SIGNAL SIGINT")
				process.Kill()
			}
		}
	}
}
