package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	values := map[string]int{
		"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
	fmt.Println(values)
	for k, v := range values {
		fmt.Println(k)
		fmt.Println(v)
	}
	// 读取本地文件。
	fd, err := os.OpenFile("simple_demo/compare_time.go", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println(fd.Name())
	fInfo, _ := fd.Stat()
	fmt.Println(fInfo.Size(), "B")
	KBSize := math.Ceil(float64(fInfo.Size()) / 1024)
	fmt.Println(KBSize, "kB")
	fd.Close()
	err2 := os.Remove("simple_demo/c3.txt")
	if err2 != nil {
		fmt.Println("Error:", err2)
		os.Exit(-1)
	}

	fmt.Println()
}
