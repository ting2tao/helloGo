package main

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var nums3 []int
	var res float64
	for _, v := range nums1 {
		nums3 = append(nums3, v)
	}
	for _, v := range nums2 {
		nums3 = append(nums3, v)
	}
	sort.Ints(nums3)
	if len(nums3) == 1 {
		res = float64(nums3[0])
	}

	if len(nums3)%2 == 0 {
		res = (float64(nums3[len(nums3)/2-1]) + float64(nums3[len(nums3)/2])) / 2
	} else {
		res = float64(nums3[len(nums3)/2])
	}

	return res
}
