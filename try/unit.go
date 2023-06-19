package main

import (
	"fmt"
	"regexp"
	"strings"
)

var reg = regexp.MustCompile(`\d|[a-zA-Z]|[一二三四五六七八九]`)

func main() {
	slice := []string{"沈阳新里程 - 二期 -8 栋-704", "沈阳新里程 - 二期 -8 栋 11 单元 -704", "沈阳新里程 - 二期 -9 栋 -12435单元 -704", "万科城·星光一分期-3#-1单元-2层-1021室", "蜀冈小学整体 - 第 6 栋-A2 单元 -11 层 -1101 室", "嘉兴公司泊樾湾北地块项目一分期-5#-1单元-4层-405室", "西宁万科城 (万灿) 一分期 -12#(高层)-1 单元 -3 层 -1032 室"}

	for _, str := range slice {

		fmt.Println(str, ExtractString(str, "栋"), ExtractString(str, "单元"))
	}
}

// ExtractString 获取某个关键词前的数字、字母、中文数字
// e.g str=阳新里程-二期-8栋二B3单元-704 s=单元 output=二B3单元
// e.g str=阳新里程-二期-8栋二B3单元-704 s=栋 output=8栋
func ExtractString(str, s string) string {
	str = strings.ReplaceAll(str, " ", "")
	res := ""
	strArr := strings.Split(str, s)

	if len(strArr) > 1 {
		strRune := []rune(strArr[0])
		for i := len(strRune) - 1; i >= 0; i-- {
			strItem := string(strRune[i])
			if reg.MatchString(strItem) {
				strItem += res
				res = strItem
			} else {
				break
			}
		}
	}
	return res + s
}
