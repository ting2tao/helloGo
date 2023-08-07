package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

//1 sync.Pool 的使用场景
//一句话总结：保存和复用临时对象，减少内存分配，降低 GC 压力。

func main() {
	marsh()
}

type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var buf, _ = json.Marshal(Student{Name: "tingting", Age: 25})

func marsh() {
	stu := stuPool.Get().(*Student)
	json.Unmarshal(buf, stu)
	stuPool.Put(stu)
	fmt.Println(stu.Name)
}

var stuPool = sync.Pool{New: func() interface{} {
	return new(Student)
}}
