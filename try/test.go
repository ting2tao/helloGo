package main

import (
	"fmt"
	"strconv"
	"time"
)

type Tes struct {
	A int `json:"a"`
}

func main() {
	//URL := "1242134"
	//fmt.Println(fmt.Sprintf("%s/%s", URL, "2352345"))
	//
	//clientIp := "118.249.212.43"
	//ipArr := strings.Split(clientIp, ".")
	////ip:= ipArr[0]*256*256*256+ipArr[0]
	//fmt.Println(ipArr)
	//
	//var s []Tes
	//
	//var w []Tes
	//var allQ []Tes
	////s=append(s,Tes{A: 1})
	////w = append(w,s...)
	//w = append(w, Tes{A: 1})
	//s = append(s, w...)
	//fmt.Println("s", s)
	//fmt.Println("w", w)
	//allQ = append(allQ, s...)
	//allQ = append(allQ, w...)
	//fmt.Println(allQ)
	//go P()
	//go P2()
	FindKey(2, 3)
	go Map2()
	time.Sleep(2000000)
	//fmt.Println(222)
}

func P() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("p")
}

func P2() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("p2")
}

func Map2() {
	m := make(map[int]interface{})
	for i := 6; i < 10; i++ {
		m[i] = strconv.Itoa(i) + "a"
	}
	for j, v := range m {
		fmt.Println("mj ", m[j])
		fmt.Println("j ", j)
		fmt.Println("v ", v)
	}
	for {
		j := 0
		j++
		if j > 10 {
			fmt.Println("j++", j)
			break
		}
	}
}
