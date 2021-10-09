package main

import "fmt"

// 2个栈实现队列

func main() {
	s := NewSList()
	s.appendTail(1)
	s.appendTail(2)
	s.appendTail(3)
	s.appendTail(4)
	s.appendTail(5)
	s.appendTail(6)
	fmt.Println(s)

	for {
		res := s.deleteHead()
		fmt.Println(res)
		if res == -1 {
			break
		}
	}
}

type SList struct {
	stackIn  []int
	stackOut []int
}

func NewSList() *SList {
	return &SList{
		stackIn:  nil,
		stackOut: nil,
	}
}

func (s *SList) appendTail(n int) {
	s.stackIn = append(s.stackIn, n)
}

func (s *SList) deleteHead() int {
	if len(s.stackOut) == 0 {
		if len(s.stackIn) == 0 {
			return -1
		} else {
			for len(s.stackIn) > 0 {
				s.stackOut = append(s.stackOut, s.stackIn[len(s.stackIn)-1])
				s.stackIn = s.stackIn[:len(s.stackIn)-1]
			}
		}
	}
	res := s.stackOut[len(s.stackOut)-1]
	s.stackOut = s.stackOut[:len(s.stackOut)-1]
	return res
}
