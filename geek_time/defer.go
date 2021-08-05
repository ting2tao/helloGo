package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println(deferDemo())
	automatic()
}
func deferDemo() int {
	defer fmt.Println("第一defer")
	defer fmt.Println("第2defer")
	//defer panic(errors.New("222"))
	//panic(errors.New("panicla"))
	fmt.Println("666")

	return 1
}

func automatic() {
	var n int32 = 10
	var num2 = n
	for {
		if atomic.CompareAndSwapInt32(&num2, 10, 0) {
			fmt.Println("The second number has gone to zero.")
			break
		}
		fmt.Println(num2)
		time.Sleep(time.Millisecond * 500)
	}
}
