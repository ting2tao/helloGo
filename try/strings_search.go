package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "史纯333涛222深圳Adc12"
	substr := "涛"
	substrrune := []rune(substr)
	fmt.Println(searchChinese(str, substr))
	fmt.Println(searchChinese2(str, substr))
	fmt.Println(searchChinese3(str, substrrune[0]))
	substr = "圳Adc1"
	substrrune = []rune(substr)
	fmt.Println(searchChinese(str, substr))
	fmt.Println(searchChinese2(str, substr))
	fmt.Println(searchChinese3(str, substrrune[0]))
}

func searchChinese(str string, substr string) bool {
	return strings.Contains(str, substr)
}

func searchChinese2(str string, substr string) bool {
	return strings.ContainsAny(str, substr)
}

func searchChinese3(str string, substr rune) bool {
	return strings.ContainsRune(str, substr)
}
