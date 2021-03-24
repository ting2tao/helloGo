package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 10
	fmt.Printf("%d %T %v \n", i, i, i)
	q()
}

type order struct {
	ordId      int
	customerId int
}

func query(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
}

func q() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	query(o)
}
