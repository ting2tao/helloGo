package main

import "fmt"

func main() {
	forRange()
	forNum()
	forRangeSlice()
}

func forRange() {
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)
}

// slice 或指针的值会变化的
func forRangeSlice() {
	numbers2 := []int{1, 2, 3, 4, 5, 6}
	numbers3 := numbers2
	maxIndex2 := len(numbers3) - 1
	for i, e := range numbers3 {
		if i == maxIndex2 {
			numbers3[0] += e
		} else {
			numbers3[i+1] += e
		}
	}
	fmt.Println(numbers3)
}

func forNum() {
	numbers1 := []int{1, 2, 3, 4, 5, 6}
	for i := range numbers1 {
		if i == 3 {
			numbers1[i] |= i
		}
	}
	fmt.Println(numbers1)
}
