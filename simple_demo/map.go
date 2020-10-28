package main

import (
	"fmt"
	"sort"
)

func main() {
	m1:=make(map[int]string)
	m1[1]="OK"
	m1[2]="OK2"
	m1[3]="OK3"
	m1[4]="OK4"
	m1[5]="OK5W"
	m2:=make(map[string]int)

	m2["ok"] = 1
	fmt.Println(m1)
	fmt.Println(m2["ok"])

	s1:=make([]int ,len(m1))
	i:=0
	for k := range m1{
		s1[i]=k
		i++
	}
	sort.Ints(s1)
	fmt.Println(s1)
}
