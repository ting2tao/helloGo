package main

func twoSum(nums []int, target int) []int {
	var arr []int
	for i, num1 := range nums {
		for j, n2 := range nums {
			if i != j && num1+n2 == target {
				arr = append(arr, i, j)
				return arr
			}
		}
	}
	return arr
}
