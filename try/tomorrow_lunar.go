package main

import (
	"encoding/json"
	"fmt"
	"github.com/Lofanmi/chinese-calendar-golang/calendar"
	"time"
)

func main() {
	for i := 0; i <= 300; i++ {
		fmt.Println(i, GetTomorrowLunar())
	}
}

var TomorrowLunarDate *time.Time
var TomorrowLunarDateString string

type LunarCalendar1 struct {
	Ganzhi struct {
		Animal     string `json:"animal"`
		Day        string `json:"day"`
		DayOrder   int    `json:"day_order"`
		Hour       string `json:"hour"`
		HourOrder  int    `json:"hour_order"`
		Month      string `json:"month"`
		MonthOrder int    `json:"month_order"`
		Year       string `json:"year"`
		YearOrder  int    `json:"year_order"`
	} `json:"ganzhi"`
	Lunar struct {
		Animal      string `json:"animal"`
		Day         int    `json:"day"`
		DayAlias    string `json:"day_alias"`
		IsLeap      bool   `json:"is_leap"`
		IsLeapMonth bool   `json:"is_leap_month"`
		LeapMonth   int    `json:"leap_month"`
		Month       int    `json:"month"`
		MonthAlias  string `json:"month_alias"`
		Year        int    `json:"year"`
		YearAlias   string `json:"year_alias"`
	} `json:"lunar"`
	Solar struct {
		Animal        string `json:"animal"`
		Constellation string `json:"constellation"`
		Day           int    `json:"day"`
		Hour          int    `json:"hour"`
		IsLeep        bool   `json:"is_leep"`
		Minute        int    `json:"minute"`
		Month         int    `json:"month"`
		Nanosecond    int    `json:"nanosecond"`
		Second        int    `json:"second"`
		WeekAlias     string `json:"week_alias"`
		WeekNumber    int    `json:"week_number"`
		Year          int    `json:"year"`
	} `json:"solar"`
}

func GetTomorrowLunar() string {
	t := time.Now().AddDate(0, 0, 1)
	if TomorrowLunarDate != nil && t.Day() == TomorrowLunarDate.Day() {
		return TomorrowLunarDateString
	}
	// 时间戳
	c := calendar.ByTimestamp(t.Unix())
	bytes, _ := c.ToJSON()
	var lunar LunarCalendar
	json.Unmarshal(bytes, &lunar)
	TomorrowLunarDateString = fmt.Sprintf("明日是%s，星期%s，农历%s%s。\n", t.Format("2006年01月02日"), lunar.Solar.WeekAlias, lunar.Lunar.MonthAlias, lunar.Lunar.DayAlias)
	TomorrowLunarDate = &t
	return TomorrowLunarDateString
}
