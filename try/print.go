package main

import (
	"fmt"
	"math"
	"sort"
)

func FormatFloat(f float64, prec int) float64 {
	multiple := math.Pow(10, float64(prec))
	return math.Floor(f*multiple) / multiple
}
func main() {

	fmt.Println(FormatFloat(1.2362353333333333333333344444444444, 2))
	a, b := 3319.2, 100.0
	fmt.Println(FormatFloat(math.Floor(a*b)/b, 3))
	fmt.Println(FormatFloat(56276*0.7, 2), 56276*0.7)
	fmt.Println(FormatFloat(10369.23623, 4))
	fmt.Println(FormatFloat(10369.23623, 5))

	s := "aasdcsdfaf"
	fmt.Println(orderlyQueue(s, 2), 2)
}

/*
*
899. 有序队列
给定一个字符串 s 和一个整数 k 。你可以从 s 的前 k 个字母中选择一个，并把它加到字符串的末尾。

返回 在应用上述步骤的任意数量的移动后，字典上最小的字符串 。

示例 1：

输入：s = "cba", k = 1
输出："acb"
解释：
在第一步中，我们将第一个字符（“c”）移动到最后，获得字符串 “bac”。
在第二步中，我们将第一个字符（“b”）移动到最后，获得最终结果 “acb”。
示例 2：

输入：s = "baaca", k = 3
输出："aaabc"
解释：
在第一步中，我们将第一个字符（“b”）移动到最后，获得字符串 “aacab”。
在第二步中，我们将第三个字符（“c”）移动到最后，获得最终结果 “aaabc”。

提示：
1 <= k <= S.length <= 1000
s 只由小写字母组成。
通过次数19,659提交次数30,85
*
*/
func orderlyQueue(s string, k int) string {
	if k == 1 {
		ans := s
		for i := 1; i < len(s); i++ {
			s = s[1:] + s[:1]
			if s < ans {
				ans = s
			}
		}
		return ans
	}
	t := []byte(s)
	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
	return string(t)
}
