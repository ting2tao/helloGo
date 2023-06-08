package main

import (
	"fmt"
	"sort"
)

func main() {
	slice1 := []int{1, 3, 5, 2, 23, 34, 23, 12, 1, 231, 231, 312, 6, 7434, 2, 24, 345, 12, 234}
	sort.Ints(slice1)
	fmt.Println(Deduplication[int](slice1))
}

func Deduplication[T int | float32 | float64 | string](arr []T) []T {
	var result []T
	mode := make(map[any]bool) //map的值不重要
	for _, v := range arr {
		if _, ok := mode[v]; !ok {
			result = append(result, v)
			mode[v] = true
		}
	}
	return result
}
