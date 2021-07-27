package main

/**
需要同时运行  hello.go global_scope.go

*/
import (
	"fmt"
	"github.com/slovty/patterns-for-public/pattern"
)

func main() {
	Goa()
}

func Goa() {
	println("Hello World")
	println("666")
	fmt.Println("30")

	fmt.Println(pattern.FullDISCOUNT)
}
