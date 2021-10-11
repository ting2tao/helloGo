package main

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	x = iota
	y
	z1 = "aaa"
	k
	j
	p = iota
)

func main() {
	fmt.Println(x, y, z1, k, j, p)

	var x interface{} = nil
	var a error = nil
	fmt.Println(x, a)
	s := make([]int, 0)
	arr := [1]int{}
	ch := make(chan struct{})
	//m1:=make(map[int]int) // 不能cap
	fmt.Println(cap(s), cap(arr), cap(ch))

	fmt.Printf("%T\n", s)
	fmt.Printf("%v\n", reflect.TypeOf(s))
}

func longestCommonPrefix(strs []string) string {
	var res string
	if len(strs[0]) == 0 {
		return res
	}
	for i, char := range strs[0] {
		count := 0
		for _, str := range strs {
			if strings.HasPrefix(str, strs[0][:i+1]) {
				count++
			}
		}
		if count != len(strs) {
			break
		}
		res = res + string(char)
	}
	return res
}
