package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println(runtime.NumCPU())
	wg2 := sync.WaitGroup{}
	wg2.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			wg2.Done()
		}(i)
	}
	wg2.Wait()
}
