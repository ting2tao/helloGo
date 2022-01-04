package main

import (
	"fmt"
	"time"
)

func main() {

	//ch := make(chan struct{}, 2)
	//
	//go func() {
	//	fmt.Println("第一个go ")
	//	ch <- struct{}{}
	//}()
	//fmt.Println("过了一会")
	//go func() {
	//	fmt.Println("第二个go ")
	//
	//	ch <- struct{}{}
	//}()
	//for i := 0; i < 2; i++ {
	//	<-ch
	//}

	chBatch()
}

// channel 练习
func chBatch() {
	now := time.Now()
	ch1 := make(chan int)
	//ch2 := make(chan int)
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s2 := []int{}
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for _, v := range s {
			fmt.Println(v)
			//fmt.Println(j)
			//取数据
			s2 = append(s2, v)
		}
		ch1 <- 0
	}()
	<-ch1
	fmt.Println(s2)
	fmt.Println(time.Since(now))
}

func ch4() {
	ch := make(chan int)
	go func(ch chan int) {
		ch <- 1
		ch <- 2
		ch <- 3
		fmt.Println("发送完毕")
	}(ch)

	timeOut := false
	for {
		select {
		case res := <-ch:
			fmt.Println("取出来了数据，", res)
		case <-time.After(time.Second * 5):
			fmt.Println("过了5s，")
			timeOut = true
		}
		if timeOut {
			break
		}
	}
	fmt.Println("完了")
}

// channel 练习
func ch3() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		//取数据
		for i := 0; i < 10; i++ {
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
