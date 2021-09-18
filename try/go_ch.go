package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	ch <- 10

	fmt.Println("发送成功")
	cha()
}

func ch1() {
	ch := make(chan int)
	go func(ch chan int) {
		ch <- 1
		ch <- 2
		ch <- 3
		fmt.Println("发送完毕")
	}(ch)

	for {
		select {
		case res := <-ch:
			fmt.Println("取出来了数据，", res)
		case time.After(time.Second * 5):

		}
	}
}

// channel 练习
func cha() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
}
