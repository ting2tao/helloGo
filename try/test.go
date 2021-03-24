package main

import (
	"fmt"
	"strings"
)

type Tes struct {
	A int `json:"a"`
}

func main() {
	URL := "1242134"
	fmt.Println(fmt.Sprintf("%s/%s", URL, "2352345"))

	clientIp := "118.249.212.43"
	ipArr := strings.Split(clientIp, ".")
	//ip:= ipArr[0]*256*256*256+ipArr[0]
	fmt.Println(ipArr)

	var s []Tes

	var w []Tes
	var allQ []Tes
	//s=append(s,Tes{A: 1})
	//w = append(w,s...)
	w = append(w, Tes{A: 1})
	s = append(s, w...)
	fmt.Println("s", s)
	fmt.Println("w", w)
	allQ = append(allQ, s...)
	allQ = append(allQ, w...)
	fmt.Println(allQ)
}
