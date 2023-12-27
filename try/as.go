package main

import (
	"fmt"
	"log"
	"math"
	"reflect"
	"sort"
	"time"
)

func main() {
	id := "431021199503083535"
	fmt.Println(ReplaceN(id, 6, 8))
	pho := "18867318997"
	fmt.Println(ReplaceN(pho, 3, 4))
	s := []string{"2", "4", "9"}
	s1 := []string{"9", "2", "4"}
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] > s1[j]
	})
	fmt.Println(reflect.DeepEqual(s, s1))

	days := []string{"20220807"}
	d, err := time.Parse("20060102", days[0])
	if err != nil {
		log.Fatalf("解释日期[%s]失败，标准的日期格式是 yyyyMMdd", days[0])
	}
	fmt.Println(d.Add(time.Hour * 4))
	fmt.Println(time.Now().Add(-time.Hour * 1))
	for i := 0; i < 99; i++ {
		AS()
	}

}

func AS() {
	cPendingIncome := math.Round(100000)
	settleAmount := cPendingIncome
	//offAmount := settleAmount
	totalAmount := math.Round(100000)
	bOffAmount := math.Round(100000)
	// 已推送金额
	if settleAmount < (totalAmount - bOffAmount) { // 2 小于账单金额时 ， 推送 待确认收入 c.PendingIncome
		settleAmount = cPendingIncome

	} else if settleAmount >= (totalAmount - bOffAmount) { // 1 当 待确认收入 大于账单金额时 ，只能推送账单金额，
		settleAmount = totalAmount - bOffAmount
		//offAmount = settleAmount + bOffAmount
	}
	//(totalAmount - bOffAmount) <
	//settleAmount = FormatFloat(settleAmount, 2)
	//offAmount = FormatFloat(offAmount, 2)
	//fmt.Println(settleAmount, offAmount)
}

func FormatFloat1(f float64, prec int) float64 {
	multiple := math.Pow(10, float64(prec))
	return math.Floor(f*multiple) / multiple
}

func ReplaceMiddleN(s string, n int) string {
	if len(s) < n {
		return s
	}

	middle := len(s) / 2
	start := middle - n/2
	end := start + n

	r := []rune(s)
	for i := start; i < end; i++ {
		r[i] = '*'
	}

	return string(r)
}

func ReplaceN(s string, start, n int) string {
	if len(s) < start+n {
		return s
	}
	end := start + n
	r := []rune(s)
	for i := start; i < end; i++ {
		r[i] = '*'
	}
	return string(r)
}
