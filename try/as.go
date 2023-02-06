package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

func main() {
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
	offAmount := settleAmount
	totalAmount := math.Round(100000)
	bOffAmount := math.Round(100000)
	// 已推送金额
	if settleAmount < (totalAmount - bOffAmount) { // 2 小于账单金额时 ， 推送 待确认收入 c.PendingIncome
		settleAmount = cPendingIncome

	} else if settleAmount >= (totalAmount - bOffAmount) { // 1 当 待确认收入 大于账单金额时 ，只能推送账单金额，
		settleAmount = totalAmount - bOffAmount
		offAmount = settleAmount + bOffAmount
	}
	//(totalAmount - bOffAmount) <
	settleAmount = FormatFloat(settleAmount, 2)
	offAmount = FormatFloat(offAmount, 2)
	fmt.Println(settleAmount, offAmount)
}

func FormatFloat(f float64, prec int) float64 {
	multiple := math.Pow(10, float64(prec))
	return math.Floor(f*multiple) / multiple
}
