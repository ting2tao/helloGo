package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	str := "PY0002SZ"

	reg, _ := regexp.Compile("\\d+")
	s := reg.FindAllString(str, -1)
	num, _ := strconv.Atoi(s[0])
	fmt.Println(num+1, s[0])
}
