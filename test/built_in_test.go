package test

import (
	"fmt"
	"testing"
	"time"
)

type Ad struct {
}

func TestNewMake(t *testing.T) {
	ad := new(Ad)

	fmt.Println(ad)

	aMap := make(map[int]string, 6)
	aMap[1] = "sss"
	fmt.Println(len(aMap), aMap)
	delete(aMap, 1)
	fmt.Println(len(aMap), aMap)
}

func TestTime(t *testing.T) {
	timeStr := "2020.01.01"
	newTime, err := time.ParseInLocation("2006.01.02", timeStr, time.Local)
	if err != nil {
		fmt.Println(err.Error())
	}
	timeStr = "2020.01.01"
	newTime, err = time.ParseInLocation("2006.01.02", timeStr, time.Local)
	if err != nil {
		fmt.Println(err.Error())
	}
	timeStr = "2020.01.01"
	newTime, err = time.ParseInLocation("2006.01.02", timeStr, time.Local)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(newTime)
}
