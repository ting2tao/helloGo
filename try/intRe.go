package main

import (
	"fmt"
)

func main() {
	n := 65535
	fmt.Println(integerReplacement(n))
	n = 3
	fmt.Println(integerReplacement(n))
}
func integerReplacement(n int) int {
	var min int
	for {
		if n == 1 {
			break
		}
		min++
		if n%2 == 0 {
			n = n / 2
			continue
		}

		if maxTime(n+1) > maxTime(n-1) {
			n = n + 1
		} else {
			n = n - 1
		}

	}
	return min
}

func maxTime(n int) int {
	var max int
	for {
		if n == 1 {
			break

		}
		n = n / 2
		max++
	}
	return max
}
