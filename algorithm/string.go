package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type People struct {
	name string `json:"name"`
}

func main() {
	var s []int

	fmt.Println(&s)

	s1 := make([]int, 0)
	fmt.Println(&s1)
	fmt.Println(reflect.DeepEqual(s, s1))

	js := `{
 "name":"11"
 }`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("people: ", p)

	str := reverseStr("12as哈哈哈d")
	fmt.Println(str)

	str1 := "234f asdasd sf"
	str2 := "234fsf"
	res := IsSameStr(str1, str2)
	fmt.Println(res)

	resStr := ReplaceStr(str1)
	fmt.Println(resStr)
}

//请实现⼀个算法，在不使⽤【额外数据结构和储存空间】的情况下，翻转⼀个给定的字
//符串(可以使⽤单个过程变量)。
//给定⼀个string，请返回⼀个string，为翻转后的字符串。保证字符串的⻓度⼩于等于
//5000。
func reverseStr(str string) string {
	fmt.Println(len(str))
	s := []rune(str)
	l := len(s)
	fmt.Println(l)
	if l > 6 {
		s = s[:6]
		l = len(s)
	}
	for i := 0; i < l/2; i++ {
		s[i], s[l-1-i] = s[l-1-i], s[i]
	}
	return string(s)
}

//判断两个给定的字符串排序后是否⼀致
//问题描述
//给定两个字符串，请编写程序，确定其中⼀个字符串的字符重新排列后，能否变成另⼀
//个字符串。 这⾥规定【⼤⼩写为不同字符】，且考虑字符串重点空格。给定⼀个string
//s1和⼀个string s2，请返回⼀个bool，代表两串是否重新排列后可相同。 保证两串的
//⻓度都⼩于等于5000。
func IsSameStr(str1, str2 string) bool {
	l1 := len([]rune(str1))
	l2 := len([]rune(str2))
	// 临界条件
	if l1 > 5000 || l2 > 5000 || l1 != l2 {
		return false
	}
	for _, v := range str1 {
		if strings.Count(str1, string(v)) != strings.Count(str2, string(v)) {
			return false
		}
	}
	return true
}

//字符串替换问题
//问题描述
//请编写⼀个⽅法，将字符串中的空格全部替换为“%20”。 假定该字符串有⾜够的空间存
//放新增的字符，并且知道字符串的真实⻓度(⼩于等于1000)，同时保证字符串由【⼤⼩
//写的英⽂字⺟组成】。 给定⼀个string为原始的串，返回替换后的string。
func ReplaceStr(str string) string {
	s := []rune(str)
	l := len(s)
	if l > 1000 {
		return ""
	}
	str = strings.ReplaceAll(str, " ", "%20")
	return str
}
