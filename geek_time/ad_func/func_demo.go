package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Printer func(contents string) (n int, err error)

func printToStd(str string) (bytesNum int, err error) {
	return fmt.Println(str)
}

func printSthToStd(str string) (bytesNum int, err error) {
	str = str + strconv.Itoa(int(time.Now().Unix()))
	return fmt.Println(str)
}
func main() {
	var p Printer
	p = printToStd
	_, _ = p("something")
	p = printSthToStd
	_, _ = p("hello")

	res, err := calculate(1, 2, add)
	if err != nil {
		fmt.Println(res, err.Error())
	}
	fmt.Println(res)
	res, err = calculate(1, 2, sub)
	if err != nil {
		fmt.Println(res, err.Error())
	}
	fmt.Println(res)
	res, err = calculate(1, 2, divide)
	if err != nil {
		fmt.Println(res, err.Error())
	}
	fmt.Println(res)
	res, err = calculate(1, 2, multiply)
	if err != nil {
		fmt.Println(res, err.Error())
	}
	fmt.Println(res)
}

type operate func(x, y int) float64

func add(x, y int) float64 {
	return float64(x + y)
}
func sub(x, y int) float64 {
	return float64(x - y)
}
func multiply(x, y int) float64 {
	return float64(x * y)
}
func divide(x, y int) float64 {
	return float64(x / y)
}
func calculate(x int, y int, op operate) (float64, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}
