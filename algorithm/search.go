package main

import "fmt"

func main() {
	res := search([]int{1, 2, 3, 4, 5, 6}, 2)
	fmt.Println(res)
	fmt.Println(search2([]int{-1, 2, 3, 4, 5, 6}, 12))
}

// 二分法查找
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}

// 循环
func search2(nums []int, target int) int {
	for index, v := range nums {
		if v == target {
			return index
		}
	}
	return -1
}
