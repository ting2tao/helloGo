package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("coinChange", coinChange([]int{2}, 3))
}

// 凑零钱
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	length := amount + 1
	for i := 0; i < length; i++ {
		if i != 0 {
			dp[i] = length
		}
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = int(math.Min(float64(dp[i]), float64(1+dp[i-coin])))
		}
	}
	if dp[amount] == length {
		return -1
	}
	return dp[amount]
}
