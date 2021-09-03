package main

import (
	"fmt"
	"unsafe"
)

type Foo struct {
}

type Bar3 struct {
	z bool  // 1
	x int32 // 4
	y *Foo  // 8
}
type Bar1 struct {
	z bool  // 1
	y *Foo  // 8
	x int32 // 4
}
type Bar2 struct {
	y *Foo  // 8
	z bool  // 1
	x int32 // 4
	a int64 // 8
}

func main() {

	var b3 Bar3
	fmt.Println(unsafe.Sizeof(b3)) // 16
	var b1 Bar1
	fmt.Println(unsafe.Sizeof(b1)) // 16
	var b2 Bar2
	fmt.Println(unsafe.Sizeof(b2))                                                                    // 16
	fmt.Println(unsafe.Alignof(b2), unsafe.Alignof(b2.x), unsafe.Alignof(b2.y), unsafe.Alignof(b2.z)) //
}
