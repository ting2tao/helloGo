package main

import (
	"fmt"
)

func main() {
COOL:
	for i := 0; i < 10; i++ {
		for {
			fmt.Print(i)
			continue COOL
		}
	}
}
