package main

import "fmt"
import "time"

func main() {
	time1 := "2020-12-18T15:38:48.822+08:00"
	time2 := "2020-12-21 09:04:25"
	//先把时间字符串格式化成相同的时间类型
	t1, err := time.Parse("2006-01-02 15:04:05", time1)
	t2, err := time.Parse("2006-01-02 15:04:05", time2)
	fmt.Println(t1)
	fmt.Println(t2)
	if err == nil && t1.Before(t2) {
		//处理逻辑
		fmt.Println("true")
	}
}
