package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	MachineID()
}

type Settings struct {
	StartTime      time.Time
	MachineID      func() (uint16, error)
	CheckMachineID func(uint16) bool
}

func MachineID() {
	sl := []int{1, 2, 3, 4, 5, 56, 6, 7, 8, 8, 89, 4, 34, 23}
	re1 := sl[:len(sl)/4]
	sl = sl[len(re1):]

	re2 := sl[:len(re1)]
	sl = sl[len(re2):]

	re3 := sl[:len(re2)]
	sl = sl[len(re3):]

	re4 := sl
	s := make([][]int, 0, 4)

	s = append(s, re1, re2, re3, re4)
	var wg sync.WaitGroup
	m := make(chan struct{}, 5) // 开5个并发

	for i, _hl := range s {
		m <- struct{}{}
		wg.Add(1)

		go func(l []int, j int) {
			defer func() {
				wg.Done()
				<-m
			}()
			PrintSlice(j, l)
		}(_hl, i)
	}

	wg.Wait()
	fmt.Println(re1, re2, re3, re4, sl)
}

func PrintSlice(name int, re []int) {
	if name == 1 {
		re[0] = 999
	}
	for i, v := range re {
		fmt.Println("name", name, "i", i, "v", v)
	}
}
