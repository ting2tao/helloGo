package main

import (
	"fmt"
	_ "fmt"
	"math"
)

func main() {
	var hello string = "Hello"
	fmt.Println(hello + "  world")
	var b, c int = 1, 2
	fmt.Println(b, c)
	var a *int
	fmt.Println(a)
	fmt.Println(b + c)
	println(math.Sqrt(float64(b)))
	//var a []int
	//fmt.Println(a)
	//var a map[string] int
	//fmt.Println(a)
	//var a chan int
	//fmt.Println(a)
	//var a func(string) int
	//fmt.Println(a)
	//var a error
	//fmt.Println(a)

	const LENGTH int = 10
	const WIDTH int = 56
	var area int
	const g, h, j = 1, false, "str" //多重赋值
	sum1 := 10
	area = LENGTH * WIDTH
	var sum = LENGTH + WIDTH + sum1
	fmt.Printf("面积为 : %d 平方米\n", area)
	fmt.Printf("%d  + %d + %d 和为 : %d ", LENGTH, WIDTH, sum1, sum)
}
