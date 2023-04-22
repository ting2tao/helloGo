package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2, 6, 9}))
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2, 4}))
	fmt.Println(findMedianSortedArrays2([]int{1, 3}, []int{2, 4}))
	fmt.Println(findMedianSortedArrays2([]int{1, 3}, []int{2, 4}))
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums3 := append(nums1, nums2...)
	sort.Slice(nums3, func(i, j int) bool {
		return nums3[i] < nums3[j]
	})
	if len(nums3)%2 == 1 {
		return float64(nums3[len(nums3)/2])
	} else {
		fmt.Println(nums3, nums3[len(nums3)/2-1], nums3[len(nums3)/2])
		return float64(nums3[len(nums3)/2-1]+nums3[len(nums3)/2]) / 2.0
	}
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	nums3 := append(nums1, nums2...)
	sort.Ints(nums3)
	if len(nums3)%2 == 1 {
		return float64(nums3[len(nums3)/2])
	} else {
		return float64(nums3[len(nums3)/2-1]+nums3[len(nums3)/2]) / 2
	}
}
