package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i1: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()

	wait2()
}

func wait2() {
	wg := sync.WaitGroup{}
	n := 2
	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
		fmt.Println("rountine 1")
	}()
	go func(n int) {
		defer wg.Done()
		time.Sleep(time.Second)
		fmt.Println("rountine 2 ")
		fmt.Println(n)
	}(n)
	wg.Wait()
	fmt.Println("total")
}
