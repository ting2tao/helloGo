package main

import "fmt"

func main() {
	var j = 10
	i := 0
	for i= 0;i<j;i++{
		fmt.Println("fio i 的值：",fibo(i))
	}
	//fmt.Println("fio i 的值：",fib(10))
}

func fib(N int) int{
	if N == 1 || N == 2 {
		return 1
	}else{
		return fib(N - 1) + fib(N - 2)
	}

}
func fibo(N int) int{
	if N <= 1  {
		return 1
	}else{
		return fibo(N - 1) + fibo(N - 2)
	}

}
