package main

import (
	"fmt"
	"hello/try/channel"
	"math/rand"
)

func main() {
	job := &channel.Job{
		Name: "taotao",
		ID:   12,
	}
	jobC := make(chan *channel.Job, 1)
	chHandler := &channel.Handler{JobCh: jobC}

	chHandler.Handle(job)

	stop := make(chan int, 1)
	stop <- 1
	chHandler = &channel.Handler{StopCh: stop}

	chHandler.Handle(job)

	channel.BufferChan()
	//channel.NoBufferChan()

	//runtime.Gosched()
	//runtime.GOMAXPROCS(5)
	//ch := make(chan string,1)
	//var count int
	//for i:=0;i<10;i++{
	//	go routine(i,ch,count)
	//}
	//fmt.Println(<-ch)
	//if count==10{
	//	close(ch)
	//}else{
	//	//fmt.Println(<-ch)
	//}
	//fmt.Println(count)
	//in := make(chan int, 1)
	//goFunc(in)
	//time.Sleep(time.Second)

	for elem := range getIntChan() {

		fmt.Printf("The element in intChan2: %v\n", elem)
	}
	selectChan()
	if v, ok := <-getIntChan(); ok {
		fmt.Println("sss", v)
	}
}

func selectChan() {

	// 准备好几个通道。
	intChannels := []chan int{
		make(chan int, 2),
		make(chan int, 2),
		make(chan int, 2),
	}
	// 随机选择一个通道，并向它发送元素值。
	index := rand.Intn(3)
	fmt.Printf("The index: %d\n", index)
	intChannels[index] <- index
	// 哪一个通道中有可取的元素值，哪个对应的分支就会被执行。
	select {
	case <-intChannels[0]:
		fmt.Println("The first candidate case is selected.")
	case <-intChannels[1]:
		fmt.Println("The second candidate case is selected.")
	case elem := <-intChannels[2]:
		fmt.Printf("The third candidate case is selected, the element is %d.\n", elem)
	default:
		fmt.Println("No candidate case is selected!")
	}
}

func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

func routine(i int, ch chan string, count int) (chan string, int) {
	fmt.Println(i)
	if i == 9 {
		ch <- "到9了"
	}
	count++
	return ch, count
}

type Ty struct {
	processedCnt int
	C            chan int
}

func goFunc(in chan int) {
	t := &Ty{
		processedCnt: 1,
		C:            make(chan int),
	}
	//in  <- 1
	t.C <- 1
	processedCnt := t.processedCnt
	go func() {
		// in for-select using ok to exit goroutine
		for {
			select {
			case x, ok := <-in:
				if !ok {
					return
				}
				fmt.Printf("Process %d\n", x)
				processedCnt++
			case <-t.C:
				fmt.Printf("Working, processedCnt = %d\n", processedCnt)
			}
		}
	}()
}
