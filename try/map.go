package main

import "fmt"

func main() {
	PointStatusText := map[string]string{
		"notDelivery": "待发货",
		"delivery":    "已发货",
		"cancel":      "已取消",
	}

	fmt.Println(PointStatusText["delivery"])
	fmt.Println(PointStatusText["notNeedDelivery"])
}
