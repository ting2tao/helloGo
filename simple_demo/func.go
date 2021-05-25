package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
	fib1()

}

// 返回一个“返回int的函数”
func fibonacci() func(int) int {
	arr := make([]int, 10)
	return func(i int) int {
		if i < 2 {
			arr[i] = i
			return i
		}
		arr[i] = arr[i-1] + arr[i-2]
		return arr[i]
	}

}

func fib1() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
