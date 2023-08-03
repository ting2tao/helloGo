package main

import (
	"fmt"
	"reflect"
)

// 定义结构体
type Person struct {
	Name    string `json:"name,omitempty" description:"55"`
	Age     int    `json:"age,omitempty"  description:"66"`
	Address string `json:"address,omitempty"  description:"77"`
}

func main() {
	// 创建一个 Person 对象
	person := &Person{Name: "Alice", Age: 20}

	// 使用 reflect 包来遍历结构体的字段
	//var fields []reflect.StructField
	s := reflect.TypeOf(person).Elem()
	s2 := reflect.ValueOf(*person)
	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		// 输出字段的值
		fmt.Printf("Field %d: %v,%v\n", i, field.Tag.Get("description"), field.Name)

	}
	for k := 0; k < s2.NumField(); k++ {
		fmt.Printf("Field %d: %v ,%v\n", k, s2.Field(k), s2.Field(k).IsValid())
		if !s2.Field(k).IsValid() {
			fmt.Println("666", s.Field(k).Tag.Get("description"))
		}
	}
}
