package main

import "fmt"

func main() {
	fmt.Println(fib(30))
	fmt.Println(dpHelper(30))
	fmt.Println(niceDpHelper(30))
	fmt.Println(leeCodeFib(30))
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
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 自下而上的dp 数组的迭代解法 优化 状态压缩 使得空间复杂度
// 一般来说是把一个二维的 DP table 压缩成一维，即把空间复杂度从 O(n^2) 压缩到 O(n)
func niceDpHelper(n int) int {
	if n < 2 {
		return n
	}
	prev := 1
	curr := 1
	for i := 3; i <= n; i++ {
		sum := prev + curr
		prev = curr
		curr = sum
	}
	return curr
}

func leeCodeFib(n int) int {
	if n < 2 {
		return n
	}
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
