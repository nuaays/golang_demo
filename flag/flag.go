package main

import (
	"flag"
	"fmt"
	//	"os"
)

func main() {
	//	for _, item := range os.Args {
	//		if item == "-d" {
	//			fmt.Println("deamon")
	//		}
	//	}

	deamon := flag.Bool("d", true, "deamon")
	fmt.Print(*deamon)
}
