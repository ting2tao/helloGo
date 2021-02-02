package main

import (
	"fmt"
)
import "time"

func main() {
	//tim:="2020-12-17 10:02:17.773827989 +0000 UTC m=+0.000087711"
	fmt.Println("2020-12-16T19:47:53.621+08:00")

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	t1, _ := time.Parse("2006-01-02 15:04:05", "2016-01-02 15:04:05")
	t2, _ := time.Parse("2006-01-02 15:04:05", "2016-01-02 15:04:05")
	fmt.Println(t1.Format("2006-01-02"))
	fmt.Println(t2.Format("2006-01-02"))

	if t1.Format("2006-01-02") == t2.Format("2006-01-02") {
		fmt.Sprintln("isToday")
	}

	startTime := time.Now().Add(time.Second * 1000)
	t3, _ := time.Parse("2006-01-02 15:04:05", startTime.Format("2006-01-02 15:04:05"))
	t4, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))

	if t3 == t4 {
		fmt.Sprintln("isToday")
	}
}
