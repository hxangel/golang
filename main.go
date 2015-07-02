package main

import (
	"fmt"
	"os/exec"
)

func main() {
	path, err := exec.LookPath("echo")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	//测试环境为ArchLinux
	fmt.Printf("echo is available at %s\n", path)    //echo is available at /usr/bin/echo
}
