package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.Version())
	runtime.Gosched()
	runtime.GOMAXPROCS(8)
}
