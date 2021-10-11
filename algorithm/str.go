package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	sArr := []rune(s)
	var stack []rune
	for _, v := range sArr {
		v1 := string(v)
		stack = append(stack, v)

		if v1 == "(" || v1 == "]" || v1 == "}" {
			temp := string(stack[len(stack)-1])

			if temp == "(" && v1 != ")" {
				return false
			}
			if temp == "[" && v1 != "]" {
				return false
			}
			if temp == "{" && v1 != "}" {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return true
}

func str2Int(s string) int {

	var n int = 0
	for i := 0; i < len(s); i++ {
		c := int(s[i])
		n = 10*n + (c - '0')
	}
	return n
}

func main() {
	s := "345还好6"
	fmt.Println(str2Int(s))
	fmt.Println(len(s))
	sres := []rune(s)

	fmt.Println(sres, len(sres))
	a := []string{"1", "2"}

	b := strings.Join(a, "")
	fmt.Println(b)
	str := "aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbcccccccccccccccccccccccccddddddddddddddeeeeeeeee"
	fmt.Println(getSameStr(str))
}
func getSameStr(str string) (string, int) {
	strArr := strings.Split(str, "")
	maxValue := strArr[0]
	maxLength := strings.Count(str, maxValue)
	strArrRes := ""
	for _, v := range strArr {
		tmpLength := maxLength
		if v != maxValue {
			tmpLength = strings.Count(str, v)
		}
		if maxLength < tmpLength {
			maxLength = tmpLength
			maxValue = v
			strArrRes = ""
		}

		if v == maxValue {
			strArrRes += v
		}
	}
	return strArrRes, maxLength
}
