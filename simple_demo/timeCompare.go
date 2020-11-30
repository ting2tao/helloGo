package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Local())
	fmt.Println("--------------")
	//
	currentTime := time.Now().Local()
	newFormat := currentTime.Format("2006-01-02 15:04:05.000")
	fmt.Println(newFormat)
	fmt.Println("--------------")
	myTime := time.Date(2018, time.December, 17, 23, 59, 59, 999, time.UTC)
	myTime1 := time.Date(2018, time.December, 17, 23, 59, 59, 999, time.UTC)
	fmt.Println("MyTime:", myTime.Format("2006-01-02 15:04:05.000"))
	fmt.Println("--------------")
	if myTime == myTime1 {
		fmt.Println("=======")
	}

	// 比较
	format := "2006-01-02 15:04:05"
	now := time.Now()
	//now, _ := time.Parse(format, time.Now().Format(format))
	a, _ := time.Parse(format, "2019-03-10 11:00:00")
	b, _ := time.Parse(format, "2015-03-10 16:00:00")

	fmt.Println("now:", now.Format(format), "\na:", a.Format(format), "\nb:", b.Format(format))
	fmt.Println("---    a > now  >  b   -----------")
	fmt.Println("now  a   After: ", now.After(a))
	fmt.Println("now  a   Before:", now.Before(a))
	fmt.Println("now  b   After:", now.After(b))
	fmt.Println("now  b   Before:", now.Before(b))
	fmt.Println("a  b   After:", a.After(b))
	fmt.Println("a  b   Before:", a.Before(b))
	fmt.Println(now.Format(format), a.Format(format), b.Format(format))
	fmt.Println(now.Unix(), a.Unix(), b.Unix())

}
