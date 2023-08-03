package main

import "fmt"

func main() {
	forScan()
}

func forScan() {
	var a, b int
	for i := 0; i < 50; i++ {
		if a/2 == 0 {
			tmpA := &a
			tmpA = &i
			tmpB := &b
			tmpB = &i

			fmt.Println(a, b, tmpA, tmpB)
		}

	}
}
