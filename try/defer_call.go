package main

import "fmt"

func main() {
	i := deferCallSelf()
	fmt.Printf("%d", i)
}
func deferCallSelf() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()

	md := make(map[int]struct{}, 10)
	for i := 0; i < 10; i++ {
		md[i] = struct{}{}
	}
	for k, _ := range md {
		md[12] = struct{}{}
		fmt.Println(md)
		delete(md, k)
		fmt.Printf("%d", len(md))
		fmt.Println(md[k])

	}
	fmt.Println(md)
	return i
}
