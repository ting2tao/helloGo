package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	//func Post(url string, bodyType string, body io.Reader) (resp *Response, err error)
	resp, err := http.Post("http://manage.xiniaodev.com:3000/mock/25/appv5/labelList",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb")) //这里的第二个参数是传入参数的类型，第三个参数固定类型为io.Reader类型，因此调用了strings包中的func NewReader(s string) *Reader 转化为io.Reader类型
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)

	decoder := json.NewDecoder(resp.Body)

	var frontReq1 Req1
	err = decoder.Decode(&frontReq1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(frontReq1)
	fmt.Println(frontReq1.Message)
	fmt.Println(frontReq1.DATA)

	//fmt.Println(string(body))
}
