package main

import (
	"fmt"
	"sync"
)

type name interface {
	PrintName()
}

var syncPool = sync.Pool{New: func() interface{} { return new(name) }}

func main() {
	i := syncPool.Get()
	fmt.Println(i)
}

func PrintName() {
	fmt.Println("sct")
}
