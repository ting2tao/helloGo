package main

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	_ "io/ioutil"
	"net/http"
	"strings"
)
type Data struct {
	LabelList  []LabelList `json:"label_list"`

}
type LabelList struct {
	Id string `json:"id"`
	Name string `json:"name"`
}
type Req struct {
	Code int
	Message string
	DATA Data
}

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
	//fmt.Println(string(body))
	decoder := json.NewDecoder(resp.Body)
	var frontReq Req
	err = decoder.Decode(&frontReq)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(frontReq)
	fmt.Println(frontReq.Message)
	fmt.Println(frontReq.DATA.LabelList[0])
	fmt.Println(frontReq.DATA.LabelList[0].Id)

	DataInfo := frontReq.DATA.LabelList

	for index ,item :=range DataInfo{
		fmt.Println("index :",index)
		fmt.Println("name :", item.Name)
	}

}
