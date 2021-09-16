package channel

import "fmt"

// 带缓冲得通道
func BufferChan() {
	ch := make(chan int, 4)
	ch <- 1
	ch <- 10
	ch <- 100
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// 不带缓冲得通道
func NoBufferChan() {
	ch := make(chan int)
	ch <- 1
	fmt.Println(<-ch)
	ch <- 10
	fmt.Println(<-ch)
	ch <- 100
	fmt.Println(<-ch)
}
