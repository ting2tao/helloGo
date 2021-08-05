package main

import (
	"fmt"
	"runtime"
	"sort"
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
	a := []int{5, 3, 1, 2, 3}
	b := []int{1, 2, 3}
	sort.Ints(a)
	sort.Ints(b)
	c := []string{"A", "B", "Hello", "D", "E", "C"}
	sort.Strings(c)
	fmt.Println(c)
	fmt.Println(a)
	fmt.Println(b)
}
