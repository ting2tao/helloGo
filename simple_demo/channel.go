package main

import "fmt"

func main() {
	// 这里我们定义了一个可以存储整数类型的带缓冲通道
	// 缓冲区大小为2
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5


	for i:=0; i<5; i++  {
		fmt.Println(<-ch)
	}
}
