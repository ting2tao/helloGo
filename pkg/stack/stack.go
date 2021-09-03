package main

import "fmt"

func main() {
	s := Init()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	fmt.Println(s)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

}

type Stack struct {
	items []interface{}
}

func Init() *Stack {
	return &Stack{items: []interface{}{}}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}
func (s *Stack) Pop() interface{} {
	if len(s.items) == 0 {
		panic("empty stack")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return item

}
