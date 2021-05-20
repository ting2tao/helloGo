package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)

	wg.Add(2)

	fmt.Println("Start........")

	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting to finish")
	wg.Wait()
	fmt.Println("Done!")
}

func printPrime(prefix string) {
	wg := sync.WaitGroup{}
	defer wg.Done()

next:
	for outer := 2; outer < 50000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s: %d\n", prefix, outer)
	}
	fmt.Printf("Finish %s\n", prefix)
}
