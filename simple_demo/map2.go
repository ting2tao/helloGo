package main

import "fmt"

func main() {
	m1:=make(map[int]string)
	m1[1]="OK"
	m1[2]="OK2"
	m1[3]="OK3"
	m1[4]="OK4"
	m1[5]="OK5W"
	m2:=make(map[string]int)

	fmt.Println(m1)
	fmt.Println(m2)

	for k,v:=range m1 {
		m2[v] =k
	}
	fmt.Println(m1)
	fmt.Println(m2)
}
