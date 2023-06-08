package main

import (
	"fmt"
	"hello/pkg/utils"
	"net/http"
)

func main() {
	response, body, errs := utils.Post("https://www.baidu.com").Send(struct {
	}{}).End()
	if errs != nil {
		fmt.Print(errs[0].Error())
	}
	if response.StatusCode != http.StatusOK {
		fmt.Print("response.StatusCode!=http.StatusOK")
	}
	fmt.Print("body：", body)
}
