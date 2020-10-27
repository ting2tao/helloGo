package main

import (
	"fmt"
	"os"
	"runtime"
)

// 查看版本
func main() {
	fmt.Println("runtime2 is " + runtime.Version())
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)

}
