package main

import (
	"fmt"
	"sync"
	"time"
)

// 多生产者消费者
var wg2 sync.WaitGroup

type PCD struct {
	Name string
	ID   int
}

func main() {
	ch := make(chan *PCD, 2)

	go product(20, ch)

	for i := 0; i < 5; i++ {
		wg2.Add(1)
		go do(ch, &wg2)
	}
	wg2.Wait()
}

func do(ch chan *PCD, wg *sync.WaitGroup) error {
	defer wg.Done()
	//for i := range ch {
	//	fmt.Println("消费", i)
	//}
	for {
		select {
		case i, ok := <-ch:
			if !ok {
				fmt.Println("channel is empty")
				return nil
			}
			fmt.Println("消费", i.Name)
		case <-time.After(time.Second * 5):
			fmt.Println("timeout")
			return nil

		}
	}

}

func product(n int, ch chan *PCD) {
	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			time.Sleep(time.Second * 2)
		}
		pcd := &PCD{
			Name: fmt.Sprintf("Name%d", i),
			ID:   i,
		}
		ch <- pcd
	}
	close(ch)
}
