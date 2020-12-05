package main

import (
	"encoding/json"
	"fmt"
)

// Product _
type Product struct {
	Name      string  `json:"name"`
	ProductID int64   `json:"product_id,omitempty,string"`
	Number    int     `json:"number,string"`
	Price     float64 `json:"price,string"`
	IsOnSale  bool    `json:"is_on_sale,string"`
}

func main() {
	var data = `{"name":"Xiao mi 6","product_id":"","number":"10000","price":"2499","is_on_sale":"false"}`
	p := &Product{}
	err := json.Unmarshal([]byte(data), p)
	fmt.Println(err)
	fmt.Println(*p)
}
