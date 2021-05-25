package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Monday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	m()
	m1()
	//fmt.Println("hello")
}

func m() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

func m1() {
	fmt.Println("counting")
	defer fmt.Println("done")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

}
