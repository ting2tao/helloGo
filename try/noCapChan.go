package main

import "fmt"

var ch chan int = make(chan int)

func foo(id int) { //id: 这个routine的标号
	ch <- id
}

func main() {
	// 开启5个routine
	for i := 0; i < 5; i++ {
		//tmp := i
		go func(j int) {
			ch <- j
		}(i)
	}

	// 取出信道中的数据
	for i := 0; i < 5; i++ {
		fmt.Print(<-ch)
	}
}
