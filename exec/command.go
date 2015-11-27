package main

import (
	//	"os"
	"os/exec"
	"syscall"
)

func main() {

	cmd := exec.Command("/bin/sh", "-c", "ls>a")
	//	cmd.Stdout = os.Stdout
	//	cmd.Stderr = os.Stderr
	cmd.Start()

	//	cmd.Run()
	//	cmd.Wait()
}
