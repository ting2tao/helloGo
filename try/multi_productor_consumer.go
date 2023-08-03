package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)

	go product(15, ch)
	close(ch)
	for i := range ch {
		fmt.Println(i)
	}

}

func product(n int, ch chan int) {
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			time.Sleep(10000)
		}
		ch <- i
	}

}
