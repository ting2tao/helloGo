package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	swg     sync.WaitGroup
	counter int
	mutex   sync.Mutex
)

func main() {

	swg.Add(2)
	go Prime(1)
	go Prime(2)
	swg.Wait()
	fmt.Println(counter)

}

func Prime(id int) {
	defer swg.Done()
	for count := 1; count < 3; count++ {
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
	}
}
