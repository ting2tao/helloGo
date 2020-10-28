package main

import "fmt"

func main() {
	//使用make 定义slice
	//s1:=make([]int,3,10)
	//fmt.Println(s1)
	//fmt.Println("元素数量：",len(s1),"容量：",cap(s1))

	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p'}
	s2 := a[2:5]
	println(&a[2])
	println(s2)
	println(string(s2))
	fmt.Println("元素数量：",len(s2),"容量：",cap(s2))

	s3:=s2[1:3]
	println(string(s3))
	fmt.Println("元素数量：",len(s3),"容量：",cap(s3))
}
