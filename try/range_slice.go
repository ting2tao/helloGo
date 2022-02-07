package main

import "fmt"

// 幸运物
var ins = []int{}

func FillIns() {
	for i := 0; i < 65; i++ {
		ins = append(ins, i)
	}
}

func main() {
	FillIns()

	waitCh := make(chan struct{})
	for i := 0; i < len(ins); i += 10 {

		end := i + 10
		if end > len(ins) {
			end = len(ins)
		}
		fmt.Println(ins[i:end], i, end)

	}
	for range waitCh {
		<-waitCh
	}
	fmt.Println(ins[100:101], len(ins))
}
