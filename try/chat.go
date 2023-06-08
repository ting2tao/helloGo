package main

import (
	"fmt"
	"sync"
	"time"
)

// 消息结构体
type Message struct {
	Content  string // 消息内容
	Sender   string // 发送者
	Receiver string // 接收者
}

func main() {
	// 创建两个线程处理消息
	go func() {
		for {
			// 等待一段时间
			time.Sleep(time.Second)

			// 处理消息
			for {
				select {
				case msg := <-messageChan:
					fmt.Println("Received message:", msg.Content, "from", msg.Sender, "to", msg.Receiver)
				}
			}
		}
	}()

	// 创建一个线程发送消息
	go func() {
		for {
			// 等待一段时间
			time.Sleep(time.Second)

			// 发送消息
			sender := "UserA"
			receiver := "UserB"
			message := Message{Content: "Hello,", Sender: sender, Receiver: receiver}
			messageChan <- message
		}
	}()

	// 启动三个线程
	fmt.Println("Started three threads")
	<-messageChan // 等待所有消息发送完毕
}

// 消息处理通道
var messageChan chan Message

// 定义全局变量
var mutexChat sync.Mutex

func init() {
	// 通道初始化为空
	messageChan = make(chan Message)
}

// 发送消息
func sendMessage(content string, sender string, receiver string) {
	mutexChat.Lock()
	defer mutexChat.Unlock()

	message := Message{Content: content, Sender: sender, Receiver: receiver}
	messageChan <- message
}
