package channel

import (
	"fmt"
	"sync"
	"time"
)

func Run() {
	startTask()
}

func do(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {

		select {
		case i, ok := <-ch:

			if ok {
				fmt.Println(i)

			} else {
				fmt.Println("no data")
				return
			}

		case <-time.After(time.Second * 5):
			fmt.Println("timeout")
			return
		}

	}

}

func startTask() {
	var wg sync.WaitGroup
	ch := make(chan int, 5)
	wg.Add(1)
	go do(ch, &wg)
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
}
