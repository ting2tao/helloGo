package main

import "fmt"

func main() {
	fmt.Println(fib(30))
}
func fib(n int) int {
	if n < 2 {
		return n
	}
	m := make(map[int]int, n)
	return helper(m, n)
}

// 带备忘录的递归
func helper(m map[int]int, n int) int {
	if n < 2 {
		return n
	}
	if m[n] != 0 {
		return m[n]
	}
	m[n] = helper(m, n-1) + helper(m, n-2)
	return m[n]
}

// 自下而上的dp 数组的迭代解法
func dpHelper(n int) int {
	if n < 2 {
		return n
	}
	dp := []int{}
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; n++ {
		dp[n] = dp[n-1] + dp[n-2]
	}
	return dp[n]
}
