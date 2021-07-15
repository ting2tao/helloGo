package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	wait1()

	wait2()

	wait3()
}

func wait1() {
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

			if i > 4 {
				time.Sleep(time.Second * 20)
			}
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
func wait2() {
	fmt.Println("wait2 start################################")
	wg := sync.WaitGroup{}
	n := 2
	wg.Add(3)
	go func() {
		defer wg.Done()
		n++
		time.Sleep(time.Second)
		fmt.Println("rountine 1")
	}()
	go func(n int) {
		defer wg.Done()
		time.Sleep(time.Second)
		fmt.Println("rountine 2 ")

	}(n)
	go func(n int) {
		defer wg.Done()
		time.Sleep(time.Second)
		fmt.Println("rountine 3 ")

	}(n)
	wg.Wait()
	fmt.Println("n:", n)
	fmt.Println("wait2 end ################################")
}

func wait3() {

	var wg sync.WaitGroup

	wg.Add(3)
	go func(n int) {
		fmt.Println("n:", n)
		t := time.Duration(n) * time.Second
		fmt.Println("t:", t)
		time.Sleep(t)

		wg.Done()
	}(1)

	go func(n int) {
		fmt.Println("n:", n)
		t := time.Duration(n) * time.Second
		fmt.Println("t:", t)
		time.Sleep(t)

		wg.Done()
	}(2)

	go func(n int) {
		fmt.Println("n:", n)
		t := time.Duration(n) * time.Second
		fmt.Println("t:", t)
		time.Sleep(t)

		wg.Done()
	}(3)

	wg.Wait()

	fmt.Println("main exit...")
}
