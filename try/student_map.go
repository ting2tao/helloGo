package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func baseStudent() {
	m := make(map[string]student)
	students := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "Hwang", Age: 22},
	}
	for _, stu := range students {
		stu.Age = 211
		m[stu.Name] = stu
	}
	fmt.Println(m)
	fmt.Println(m["li"].Age)
}

func main() {
	baseStudent()
}
