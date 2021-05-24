package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/cmplx"
	"regexp"
	"strings"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	fmt.Println(sqrt(2), sqrt(-4))

	arr()
	slice()

	slice2()
	lenAndCap()
	slice3()
	makeSlice()
	slice2slice()

	fmt.Println(FindDigits("D:/GoWorkSpace/src/hello/try/c3.txt"))
	findStr := fmt.Sprintf("%x", FindDigits("D:/GoWorkSpace/src/hello/try/c3.txt"))
	fmt.Println(findStr)
}

func slice2slice() {
	// 创建一个井字板（经典游戏）
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func makeSlice() {
	b := make([]int, 5, 5) // len(b)=0, cap(b)=5
	fmt.Printf("len=%d cap=%d %v\n",
		len(b), cap(b), b)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func arr() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slice() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	names[0] = "ShiChuntao"
	fmt.Println(a, b)
	fmt.Println(names)
}

func slice2() {
	q := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	s1 := s[4:]
	fmt.Println(s)
	fmt.Println(s1)
}

func lenAndCap() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func slice3() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

var digitRegexp = regexp.MustCompile("hello world")

func FindDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	var c []byte
	c = append(c, b...)
	//fmt.Printf("%q",c)
	return c

}
