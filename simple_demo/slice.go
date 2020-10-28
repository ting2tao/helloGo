package main

import "fmt"

func main() {
	// slice 的定义
	var s1 []int
	fmt.Println(s1)
	a := [...]int{1,2,3,5,6,7,8,9,10,11,12,13}
	fmt.Println(a)
	arr := [10]int{}
	fmt.Println(arr)
	s2 := a[5:]
	fmt.Println(s2)
}
