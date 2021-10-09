package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	s := "cat"
	runtime.GOMAXPROCS(1)
	for j := 0; j < 2; j++ {
		go printStr(fmt.Sprintf("%s%d", s, j))
	}
	s = "dog"
	for j := 0; j < 2; j++ {
		go printStr(fmt.Sprintf("%s%d", s, j))
	}
	s = "fish"
	for j := 0; j < 2; j++ {
		go printStr(fmt.Sprintf("%s%d", s, j))
	}
	time.Sleep(time.Second)
}

func printStr(s string) {
	fmt.Println(s)
}
