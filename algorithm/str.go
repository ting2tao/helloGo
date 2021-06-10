package main

import (
	"fmt"
	"strings"
)

func main() {
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
