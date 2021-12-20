package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

func main() {
	//fmt.Println(time.Now().GoString())
	//TestM()
	now := time.Now()
	//time.Sleep(3 * time.Second)
	//fmt.Println(time.Since(now))
	//
	timeObj := time.Unix(now.UnixMilli()/1000, 0)
	fmt.Println(timeObj.Hour())
	//
	//mobiles := []string{"188", "122"}
	//fmt.Println("('" + strings.Join(mobiles, "','") + "')")
	//hours := []int{6, 23, 20, 21, 2, 5, 4, 5}
	//sort.Ints(hours)
	//fmt.Println(hours)
	//for i := 0; i < 10; i++ {
	//	fmt.Println(getMaxCountHour(hours))
	//}

	GetWeekPeriodDate()
	//WeekDayTest()
}

func GetWeekPeriodDate() {
	now := time.Now().AddDate(0, 0, 11)
	fmt.Println(now)
	offset := int(7 - now.Weekday())
	if offset == 7 {
		offset = 0
	}
	weekEnd := time.Date(now.Year(), now.Month(), now.Day(), 20, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	fmt.Println(weekEnd)
	fmt.Println(weekEnd.UnixMilli())
	weekStart := weekEnd.AddDate(0, 0, -7)
	fmt.Println(weekStart)
	fmt.Println(weekStart.UnixMilli())
	for i := 2; i < 10; i++ {
		weekStart = weekEnd.AddDate(0, 0, -7*i)
		fmt.Println(weekStart)
		fmt.Println(weekStart.UnixMilli())
		fmt.Println(weekStart.Unix())
	}

	//uni := weekEnd.UnixMilli()
	//timeobj := time.UnixMilli(uni)
	//
	//date := timeobj.Format("2006-01-02 15:04:05")
	//
	//fmt.Println(date)
	//timeobj = time.UnixMilli(1639324850000)
	//date = timeobj.Format("2006-01-02 15:04:05")
	//fmt.Println(date)
	//fmt.Println(timeobj.Hour())
}

type RuiYiXingMQMsg struct {
	Action      int64  `json:"action"`
	External    string `json:"external"`
	Key         string `json:"key"`
	ProjectCode string `json:"projectCode"`
}

type External struct {
	DeviceCode string `db:"device_code" json:"deviceCode"` // 设备编码
	DeviceName string `db:"device_name" json:"deviceName"` // 设备名称
	PassType   int64  `db:"pass_type" json:"passType"`     //: 通行方式: 10刷卡 20二维码  30人脸  40蓝牙 50 远程  60 对讲
	CardNo     string `db:"card_no" json:"cardNo"`         // 卡号
	UserCode   string `db:"user_code" json:"userCode"`     //用户编码（睿易行）
	FaceId     string `db:"face_id" json:"faceId"`         // 抓拍图
	QrCode     string `db:"qr_code" json:"qrCode"`         //二维码内容
	PassTime   string `db:"pass_time" json:"passTime"`     //通行时间
	Direction  int64  `db:"direction" json:"direction"`    //0 不限定; 1 入口； 2 出口
	IdCard     string `db:"id_card" json:"idCard"`         //  身份证号码
	Username   string `db:"username" json:"username"`      //用户姓名

	ProjectCode string `db:"project_code" json:"projectCode"` //项目编码
	Telephone   string `db:"telephone" json:"telephone"`      // 手机号

	OriginalContent string    `db:"original_content"` // 用于记录原始消息信息
	CreatedAt       time.Time `db:"created_at"`
}

func TestM() {
	var r RuiYiXingMQMsg
	var ex External
	str := `{"action":1,"external":"{\"cardNo\":\"FA63D276\",\"deviceCode\":\"3100000000001200\",\"direction\":0,\"idCard\":\"\",\"passTime\":\"1639448584000\",\"passType\":10,\"projectCode\":\"44030145\",\"sourceCode\":\"91059527340852428801\",\"sourceType\":\"SX\",\"telephone\":\"13168758723\",\"userCode\":\"126975949\",\"userName\":\"森林\"}","key":"PASS_","projectCode":"44030145"}`
	json.Unmarshal([]byte(str), &r)
	json.Unmarshal([]byte(r.External), &ex)
	fmt.Println(r)
	fmt.Println(ex)
}

func getMaxCountHour(hours []int) int {
	sort.Ints(hours)
	temp := make(map[int]int, len(hours))
	for _, n := range hours {
		temp[n]++
	}
	maxV, maxK := 0, 0
	for hour, count := range temp {
		if maxV < count {
			maxV, maxK = count, hour
			continue
		}
		if maxV == count {
			if hour < 5 { //5点之前算晚上
				if maxK >= 5 {
					maxK = hour
					continue
				}
				if maxK < hour {
					maxK = hour
					continue
				}
				continue
			}

			if maxK < hour && maxK >= 5 {
				maxK = hour
			}

		}
	}
	return maxK
}

func WeekDayTest() {
	now := time.Now().AddDate(0, 0, 4)
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	fmt.Println(weekStart)
	fmt.Println(weekStart.UnixMilli())

	offset = int(7 - now.Weekday())
	if offset == 7 {
		offset = 0
	}
	weekEnd := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.Local).AddDate(0, 0, offset)
	fmt.Println(weekEnd)
	fmt.Println(weekEnd.UnixMilli())

	uni := weekEnd.UnixMilli()
	timeobj := time.UnixMilli(uni)

	date := timeobj.Format("2006-01-02 15:04:05")

	fmt.Println(date)
	timeobj = time.UnixMilli(1639324850000)
	date = timeobj.Format("2006-01-02 15:04:05")
	fmt.Println(date)
	fmt.Println(timeobj.Hour())
}
