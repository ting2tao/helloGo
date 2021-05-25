package main

import "fmt"

func main() {
	m()
	printStruct()
}

type Vertex struct {
	X int
	Y int
}

func printStruct() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

func m() {
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值 通过指针已经把底层的值改掉
	fmt.Println(*p) // 查看 *p 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
}
