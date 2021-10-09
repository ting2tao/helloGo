package main

import (
	"fmt"
	"math/bits"
)

func main() {
	//r := NewRPCError(4004,"sss")
	//fmt.Println(AsErr(r))
	a1 := countBits(5)
	fmt.Println(a1)
}

func countBits(n int) []int {
	a := []int{}
	for i := 0; i <= n; i++ {
		a = append(a, bits.OnesCount(uint(i)))
	}
	return a
}

type RPCError struct {
	Code int
	Msg  string
}

func (r *RPCError) Error() string {
	return fmt.Sprintf("%s,code:%d", r.Msg, r.Code)
}

func NewRPCError(code int, msg string) error {
	return &RPCError{
		Code: code,
		Msg:  msg,
	}
}

func AsErr(err error) error {
	return err
}

type bird interface {
	Fly() string
	SayHi() string
}
