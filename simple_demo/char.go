package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "1_ABD_129_0"
	str6 := "ABCDF"
	fmt.Println(string(str6[0]))
	var s1 []string
	s1 = strings.Split(str, "_")
	fmt.Println(s1[0]) //id
	fmt.Println(s1[1]) //选项
	fmt.Println(s1[2]) //用时
	fmt.Println(s1[3]) //是否标记

	var s2 []string
	s2 = strings.Split(s1[1], "")
	fmt.Println(s2) //是否标记

	fmt.Println(strings.Split(str, "_")[0])
}
