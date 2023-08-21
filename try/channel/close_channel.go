package channel

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func Run() {
	startTask()
}

func do(ch chan int, wg *sync.WaitGroup) {
	//defer wg.Done()
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

	go do(ch, &wg)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		ch <- i
	}
	close(ch)
	wg.Wait()
}

func TestDo(t *testing.T) {
	t.Log(runtime.NumGoroutine())
	startTask()
	time.Sleep(time.Second)
	t.Log(runtime.NumGoroutine())
}
