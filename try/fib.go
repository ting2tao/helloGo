package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(fib(30))
	fmt.Println(dpHelper(30))
	fmt.Println("niceDpHelper(30)", niceDpHelper(30))
	fmt.Println(coinChange([]int{2}, 3))
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
	dp[0], dp[1], dp[2] = 0, 1, 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 自下而上的dp 数组的迭代解法
// 缩小空间复杂度 状态压缩
func niceDpHelper(n int) int {
	if n < 2 {
		return n
	}
	sum, prev, curr := 0, 1, 1
	for i := 3; i <= n; i++ {
		sum = prev + curr
		prev = curr
		curr = sum
	}
	return curr
}

// 1 确定base case 0 返回0 小于0 返回 负1
//1、确定 base case，这个很简单，显然目标金额 amount 为 0 时算法返回 0，因为不需要任何硬币就已经凑出目标金额了。
//
//2、确定「状态」，也就是原问题和子问题中会变化的变量。由于硬币数量无限，硬币的面额也是题目给定的，只有目标金额会不断地向 base case 靠近，所以唯一的「状态」就是目标金额 amount。
//
//3、确定「选择」，也就是导致「状态」产生变化的行为。目标金额为什么变化呢，因为你在选择硬币，你每选择一枚硬币，就相当于减少了目标金额。所以说所有硬币的面值，就是你的「选择」。
//
//4、明确 dp 函数/数组的定义。我们这里讲的是自顶向下的解法，所以会有一个递归的 dp 函数，一般来说函数的参数就是状态转移中会变化的量，也就是上面说到的「状态」；
//函数的返回值就是题目要求我们计算的量。就本题来说，状态只有一个，即「目标金额」，题目要求我们计算凑出目标金额所需的最少硬币数量。所以我们可以这样定义 dp 函数：
//dp(n) 的定义：输入一个目标金额 n，返回凑出目标金额 n 的最少硬币数量。1、确定 base case，这个很简单，显然目标金额 amount 为 0 时算法返回 0，因为不需要任何硬币就已经凑出目标金额了。

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	dp := make([]int, amount+1)
	// base case
	dp[0] = 0
	// 外层 for 循环在遍历所有状态的所有取值
	for j := 0; j < len(dp); j++ {
		if j != 0 {
			dp[j] = amount + 1
		}
		// 内层 for 循环在求所有选择的最小值
		for _, coin := range coins {
			// 无解跳过
			if j-coin < 0 {
				continue
			}
			dp[j] = int(math.Min(float64(dp[j]), float64(1+dp[j-coin])))
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
