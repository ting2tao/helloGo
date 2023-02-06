package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"math"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unicode"
)

func main() {

	city := "上市区海市"
	fmt.Println(strings.TrimRight(city, "市"))

	mobileRegexp := regexp.MustCompile(`^1\d{10}$`)
	if mobileRegexp.MatchString("13538802261") {
		fmt.Println(true, time.Now().Format("20060102150405"))
	}
	fmt.Println(math.Ceil(4.0 / 3))

	t, _ := time.ParseInLocation("20060102", "20221231", time.Local)
	startDate := t.Format("200601")
	endDate := AddDate(t, 0, 1, 0).Format("200601")

	fmt.Println(startDate, endDate)

	totalAccountAmount := decimal.NewFromFloat(838280.6865)
	finalPendingAmount := decimal.NewFromFloat(2099.74455566)
	fmt.Println(totalAccountAmount.Round(2), finalPendingAmount.Round(2), reflect.TypeOf(finalPendingAmount))

	ctx := gin.Context{}
	s := struct {
		UserID   int    `json:"user_id"`
		UserName string `json:"user_name"`
	}{
		1,
		"sss",
	}
	m, _ := json.Marshal(s)
	ctx.Set("u", string(m))
	fmt.Println(ctx.GetString("u"))
	fmt.Println(StringIsDigit(""))
	fmt.Println(StringIsDigit("adasq2222"))
	fmt.Println(StringIsDigit("12112awad"))
	fmt.Println(StringIsDigit("12112"))
}

func StringIsDigit(s string) bool {
	return StringIsX(s, func(r rune) bool {
		return r < unicode.MaxLatin1 && unicode.IsDigit(r)
	})
}

type operator func(rune) bool

func StringIsX(str string, o operator) bool {
	if str == "" { // 空字符串，什么都不是。。。
		return false
	}
	for _, r := range []rune(str) {
		if !o(r) {
			return false
		}
	}
	return true
}

func AddDate(t time.Time, year, month, day int) time.Time {
	//先目标月的1号
	targetDate := t.AddDate(year, month, -t.Day()+1) // 获目标月的临界值
	targetDay := targetDate.AddDate(0, 1, -1).Day()  //对比临界值与振日期值，取最小的值
	if targetDay > t.Day() {
		targetDay = t.Day()
	}
	//最后用目标月的1号加上目标值和入参的天数
	targetDate = targetDate.AddDate(0, 0, targetDay-1+day)
	return targetDate
}
