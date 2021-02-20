package main

import (
	"fmt"
)

func a() {
	// the expression "i" is evaluated when the Println call is deferred. The deferred call will print "0" after the function returns.
	i := 0
	defer fmt.Println(i) //  A deferred function's arguments are evaluated when the defer statement is evaluated.
	i++
	fmt.Println(i)
	return
}

// 先进后出
func b() {
	// Deferred function calls are executed in Last In First Out order after the surrounding function returns.
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

func c() (i int) {
	// Deferred functions may read and assign to the returning function's named return values.
	defer func() { i++ }()
	return 1
}

func dMy() (j int, otherArr []int) {
	myArr := []int{1, 2, 3, 4}

	defer fmt.Println(otherArr) // output []
	defer fmt.Println(j)        // out put 0
	for _, i := range myArr {
		defer func(i int) {
			otherArr = append(otherArr, i)
			j++
		}(i)
	}
	return
}
func main() {
	fmt.Println("hello")
	a()
	b()
	fmt.Println(c())
	fmt.Println(dMy())
}
