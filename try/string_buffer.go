package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	var buffer bytes.Buffer
	for i := 0; i < 10; i++ {
		buffer.WriteString("a")
	}
	fmt.Println(buffer.String())
	string2()
}

//源码对于Buffer的定义中,并没有关于锁的字段,在write和read函数中也未发现锁的踪影,所以符合上面提到的文档中的rule,即Buffer并发是不安全的。
func string2() {
	var buffer bytes.Buffer
	buffer.WriteString("a")
	buffer.WriteString("4")
	fmt.Println(buffer.String())
	var bf Buffer
	i, _ := bf.Write([]byte{1, 2, 3})
	fmt.Println(i)
	fmt.Println(bf.Read([]byte{1, 2, 3}))

}

//两种锁的区别
//	sync.Mutex(互斥锁)
//当一个goroutine访问的时候，其他goroutine都不能访问，保证了资源的同步，避免了竞争，不过也降低了性能
//	sync.RWMutex(读写锁)
//非写状态时:多个Goroutine可以同时读,一个Goroutine写的时候,其它Goroutine不能读也不能写,性能好
type Buffer struct {
	b  bytes.Buffer
	rw sync.RWMutex
}

func (b *Buffer) Read(p []byte) (n int, err error) {
	b.rw.RLock()
	defer b.rw.RUnlock()
	return b.b.Read(p)
}
func (b *Buffer) Write(p []byte) (n int, err error) {
	b.rw.Lock()
	defer b.rw.Unlock()
	return b.b.Write(p)
}
func (b *Buffer) String() string {
	return b.String()
}
