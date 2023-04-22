package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Ino struct {
	RequestInfos  string `json:"request_infos"`
	RequestMethod string
}

func generateRandomString(size int) string {
	arr1 := []int{1, 2, 3, 4, 5, size}
	arr2 := []int{1, 2, 5, size, 6, 7}
	fmt.Println(arr1[1:len(arr2)], arr1[len(arr2):])
	return ""
}

func main() {
	fmt.Println(generateRandomString(16))
	crDate, err := time.ParseInLocation(time.RFC3339, "2023-03-14T04:30:27.000+00:00", time.Local)
	fmt.Println(crDate, err)
	SwitchToDate := time.Date(2023, 4, 13, 15, 29, 59, 0, time.Local)
	if time.Now().After(SwitchToDate) {
		fmt.Println("13之后")
	} else {
		fmt.Println("小于")
	}
	str := "你好，以下租客合同Z20221017115721249已退租，请及时处理退租账单：房源地址：武汉市武汉香港路8号1号楼-1单元-19楼-1918室，合同起止日期：2022.10.20 ~ 2023.01.10,退租日期：2023.03.15，操作人：杨思,退租账单：付款2200.00元，退租账单收/付款日：2023.04.14。"
	runeStr := []rune(str)
	fmt.Println(len(runeStr))
	str1 := "尊敬的客户，您在朴邻交易的郑州万科美景魅力之城幽兰园小区1栋-1-106仍未完成评价，为了给您提供更优质的服务，请您尽快完成本次服务评价，点击 https://vk8.co/4WbnTx9H3O 进入评价。签约结束，服务不止，有事帮忙找管家，祝您入住顺利！回复TD拒收"
	runeStr1 := []rune(str1)
	fmt.Println(len(runeStr1))
	str23 := "https://vk8.co/AfgiYPtcpI"
	runeStr2 := []rune(str23)
	fmt.Println(len(runeStr2))
	fmt.Println(time.Now().UnixMilli())
	fmt.Println(InnerFunc())
}

func InnerFunc() error {
	ino := Ino{
		RequestInfos:  "",
		RequestMethod: "AsyncKsbActivityOrder",
	}
	var err error
	check := func(p interface{}) error {
		err = json.Unmarshal([]byte(ino.RequestInfos), &p)
		return err
	}
	switch ino.RequestMethod {
	case "AsyncKsbActivityOrder":
		var params string
		err = check(params)
	}
	return err
}
