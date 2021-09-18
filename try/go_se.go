package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var m1 sync.Mutex

func main() {
	chan2()
}

func goWait() {
	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			out <- rand.Intn(5)
		}
		close(out)
	}()
	go func() {
		defer wg.Done()
		for i := range out {
			fmt.Println(i)
		}
	}()
	wg.Wait()
}

func chan1() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	//runtime.GOMAXPROCS(1)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(muTex(i))
		}(i)
	}
	wg.Wait()
}

func chan2() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	//runtime.GOMAXPROCS(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			if i < 3 {
				fmt.Println(muTex(i))
			}
		}
	}()
	wg.Wait()
}
func muTex(a int) int {
	m1.Lock()
	defer m1.Unlock()
	a++
	return a
}
