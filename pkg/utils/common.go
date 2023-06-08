package utils

import (
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	Digits    = "0123456789"
	UpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowerCase = "abcdefghijklmnopqrstuvwxyz"
	Specials  = "~=+%^*/()[]{}/!@#$?|"
)

var (
	location *time.Location
)

func init() {
	rand.Seed(time.Now().UnixNano())
	location, _ = time.LoadLocation("Asia/Shanghai")
}

func Fail(c *gin.Context, httpCode, bizCode int, format string, a ...interface{}) {
	c.AbortWithStatusJSON(httpCode, gin.H{
		"code": bizCode,
		"msg":  fmt.Sprintf(format, a...),
		"data": gin.H{},
	})
}

func FailAs400(c *gin.Context, format string, a ...interface{}) {
	Fail(c, http.StatusBadRequest, http.StatusBadRequest, format, a...)
}

func FailAs401(c *gin.Context) {
	Fail(c, http.StatusUnauthorized, http.StatusUnauthorized, "未授权")
}

func FailAs403(c *gin.Context, format string, a ...interface{}) {
	Fail(c, http.StatusForbidden, http.StatusForbidden, format, a...)
}

func FailAs404(c *gin.Context, format string, a ...interface{}) {
	Fail(c, http.StatusNotFound, http.StatusNotFound, format, a...)
}

func FailAs500(c *gin.Context, format string, a ...interface{}) {
	Fail(c, http.StatusInternalServerError, http.StatusInternalServerError, format, a...)
}

func Success(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "",
		"data": result,
	})
}

// RandomString 生成随机的字符串
func RandomString(table string, length int) string {
	buf := make([]byte, length)
	tableLen := len(table)
	for i := 0; i < length; i++ {
		buf[i] = table[rand.Intn(tableLen)]
	}
	return string(buf)
}

// ValidateMobile 验证提供的手机号是否合法,验证方式比较简单
func ValidateMobile(mobile string) bool {
	matched, _ := regexp.MatchString("^1[0-9]{10}$", mobile)
	return matched
}

// 校验用户身份证
func ValidateIdentityNumber(number string) bool {
	matched, _ := regexp.MatchString(`[1-6]\d{5}((19\d{2})|(20((0\d)|(1[0-8]))))((0[1-9])|(1[012]))((0[1-9])|([12]\d)|(3[01]))\d{3}[\dX]`, number)
	return matched
}

func SecurityCodeFaker() string {
	return RandomString(Digits, 6)
}

func GenerateDayKey() string {
	return time.Now().Format("20060102")
}

func ContainsUint64(list []uint64, v uint64) bool {
	for _, item := range list {
		if item == v {
			return true
		}
	}
	return false
}

func ContainsString(list []string, v string) bool {
	for _, item := range list {
		if item == v {
			return true
		}
	}
	return false
}

// AddTZ 用于把数据库中没带时区（或者说时区是 UTC）的时间，加上中国时区
func AddTZ(t time.Time) time.Time {
	if t.IsZero() {
		return t
	}
	return time.Date(
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), location,
	)
}

func AfterOrEqual(t1, t2 time.Time) bool {
	if t1 == t2 {
		return true
	}
	return t1.After(t2)
}

func BeforeOrEqual(t1, t2 time.Time) bool {
	if t1 == t2 {
		return true
	}
	return t1.Before(t2)
}

func ParseInLocation(layout string, value string) (time.Time, error) {
	return time.ParseInLocation(layout, value, location)
}

func GetWeekStartAndEndDate() (time.Time, time.Time) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)

	offset = int(7 - now.Weekday())
	if offset == 7 {
		offset = 0
	}
	weekEnd := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.Local).AddDate(0, 0, offset)

	return weekStart, weekEnd
}

// GetWeekPeriodDate 从周日到上周日
// 以周为单位，每周日八点，到下一个周日的八点为一个统计周期。八点之后的统计到下一个周期里面
func GetWeekPeriodDate(hour int) (time.Time, time.Time) {
	now := time.Now()
	offset := int(7 - now.Weekday())
	if offset == 7 {
		offset = 0
	}
	weekEnd := time.Date(now.Year(), now.Month(), now.Day(), hour, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekStart := weekEnd.AddDate(0, 0, -7)
	return weekStart, weekEnd
}

// GetPeriodTime
// 获取区间段的时间点
func GetPeriodTime(now time.Time, hour int, span int) (time.Time, time.Time) {
	end := time.Date(now.Year(), now.Month(), now.Day(), hour, 0, 0, 0, time.Local)
	start := end.AddDate(0, 0, -span)
	return start, end
}

func GenAppID() string {
	return "wj" + RandomString(Digits+LowerCase, 14)
}
