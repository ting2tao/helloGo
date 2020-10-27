package main

import (
	"fmt"
	"runtime"
)

func main() {
	//compare(20, 30)
	prompt := "Enter a digit, e.g. 3 "+ "or %s to quit."
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { //Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
	fmt.Println(prompt)
	fmt.Println(isGreater(5,2))

}


func isGreater(x, y int) string {
	if x > y {
		return "大于"
	}
	return "下雨i"
}
//func compare(num1, num2) {
//
//	if num1 > num2 {
//		fmt.Println(num1 + "大于 " + num2)
//	} else {
//		fmt.Println(num1 + "小于 " + num2)
//	}
//}
