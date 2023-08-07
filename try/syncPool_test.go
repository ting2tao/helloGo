package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func BenchmarkMarsh(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &Student{}
		json.Unmarshal(buf, stu)
	}
}
func BenchmarkMarshWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := stuPool.Get().(*Student)
		json.Unmarshal(buf, stu)
		stuPool.Put(stu)
		fmt.Println(stu.Name)
	}
}
