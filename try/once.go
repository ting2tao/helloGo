package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

//func main(){
//	//channelF()
//	sumFib()
//}

func main() {

}

func wgs() {
	ch := make(chan int, 1)
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	cou := make(chan int, 1)
	ch <- 1
	count := 0
	for i := 0; i < 100; i++ {
		go func() {
			<-ch
			fmt.Println("dog")
			ch1 <- 1
		}()
		go func() {
			<-ch1
			fmt.Println("cat")
			ch2 <- 1
		}()
		go func() {
			<-ch2
			fmt.Println("fish")
			cou <- 1
			count++
			ch <- 1
		}()
	}
	for i := 0; i < 100; i++ {
		<-cou
	}
	fmt.Println(count)
}
func wg1() {
	ch := make(chan int, 1)
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	var wg sync.WaitGroup
	ch <- 1
	count := 0
	for i := 0; i < 100; i++ {
		wg.Add(3)
		go func() {
			defer wg.Done()
			<-ch
			fmt.Println("dog")
			ch1 <- 1
		}()
		go func() {
			defer wg.Done()
			<-ch1
			fmt.Println("cat")
			ch2 <- 1
		}()
		go func() {
			defer wg.Done()
			<-ch2
			fmt.Println("fish")
			count++
			ch <- 1
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func fm() {
	ch := make(chan int, 1)
	//ch1 := make(chan int, 1)
	//ch2 := make(chan int, 1)
	var wg sync.WaitGroup
	ch <- 1
	count := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			<-ch
			ch <- 1
			fmt.Println(k)

			count++
		}(i)

	}
	wg.Wait()
	fmt.Println(count)
}

func printOdd(ch chan byte, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 9; i += 3 {
		fmt.Println(1)
		ch <- 1
		<-ch
	}
}

func printEven(ch chan byte, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 3 {

		<-ch
		fmt.Println(2)
		ch <- 3
	}
}

func printThree(ch chan byte, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 3; i <= 10; i += 3 {
		<-ch
		fmt.Println(3)
		ch <- 1
	}
}

type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) DO(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()

	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

func channelF() {
	var ch chan int
	defer close(ch)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		wg.Done()
		ch = make(chan int, 1)
		ch <- 55
	}()

	go func(ch *chan int) {
		wg.Done()
		time.Sleep(time.Second)
		var ok, v = <-*ch
		fmt.Println(ok, v)
	}(&ch)
	wg.Wait()

	c := time.Tick(time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}
func ch2() {
	ch := make(chan int)
	var count int
	go func() {
		ch <- 1
	}()
	go func() {
		count++

	}()
	<-ch
	close(ch)
	fmt.Println(count)
}

// 按1 2 3 的顺序打印100次
func chP() {

	for {
		go func() {

			fmt.Println(1)
		}()

		go func() {
			fmt.Println(2)
		}()
		go func() {
			fmt.Println(3)
		}()
	}

}

func fib2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		//case <-quit:
		//	fmt.Println("quit")
		//	return
		case <-time.After(time.Nanosecond * 1):
			fmt.Println("timeout 5s")
			return
		}
	}
}
func sumFib() {
	c, quit := make(chan int), make(chan int)
	n := 10
	go func(n int) {
		for i := 0; i < n; i++ {

			fmt.Println(<-c)

		}
		quit <- 0
	}(n)
	fib2(c, quit)
}
