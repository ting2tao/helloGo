package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 5, 5, 5, 6, 7}
	//nums :=[]int{}
	target := 5
	arrRange := searchRange(nums, target)
	fmt.Println(arrRange)
	arrRange = binarySearchRange(nums, target)
	fmt.Println(arrRange)
	arrRange = goSearchRange(nums, target)
	fmt.Println(arrRange)

	arrRange = binSearchRange(nums, target)
	fmt.Println(arrRange)
}

// o(n)
func searchRange(nums []int, target int) []int {
	arrRange := []int{-1, -1}
	for i, v := range nums {
		if v == target && arrRange[0] == -1 {
			arrRange[0] = i
			arrRange[1] = i
		}
		if v == target && i != arrRange[0] {
			arrRange[1] = i
		}
	}
	return arrRange
}

// o(logn)
func binarySearchRange(nums []int, target int) []int {
	arrRange := []int{-1, -1}
	//
	return arrRange
}

func goSearchRange(nums []int, target int) []int {
	arrRange := []int{-1, -1}
	leftBound := sort.SearchInts(nums, target)
	if leftBound == len(nums) || nums[leftBound] != target {
		return arrRange
	}
	rightBound := sort.SearchInts(nums, target+1) - 1
	return []int{leftBound, rightBound}
}

func binSearch(nums []int, target int, lower bool) int {
	left, right, idx := 0, len(nums)-1, len(nums)
	for left <= right {
		mid := int(math.Floor(float64((left + right) / 2)))
		if nums[mid] > target || (lower && nums[mid] >= target) {
			right = mid - 1
			idx = mid
		} else {
			left = mid + 1
		}
	}
	return idx
}

func binSearchRange(nums []int, target int) []int {
	arrRange := []int{-1, -1}
	leftIdx := binSearch(nums, target, true)
	rightIdx := binSearch(nums, target, false) - 1
	if leftIdx <= rightIdx && rightIdx < len(nums) && nums[leftIdx] == target && nums[rightIdx] == target {
		arrRange = []int{leftIdx, rightIdx}
	}
	return arrRange
}
