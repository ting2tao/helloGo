package cron

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hello/pkg/logger"
	"log"
	"net/http"
	"testing"
)

func TestCron(t *testing.T) {
	crontab := NewCrontab()

	// 添加函数作为定时任务
	taskFunc := func() {

		// 定义一个结构体来匹配你的JSON数据
		type ExampleRequest struct {
			Field1 string `json:"field1"`
			Field2 int    `json:"field2"`
		}

		// 创建一个ExampleRequest对象，并填充数据
		requestData := ExampleRequest{
			Field1: "value1",
			Field2: 42,
		}

		// 将结构体序列化为JSON格式的字节切片
		jsonData, err := json.Marshal(requestData)
		if err != nil {
			log.Fatalf("序列化JSON时出错: %v", err)
		}

		// 使用bytes.NewBuffer创建一个io.Reader
		reader := bytes.NewBuffer(jsonData)

		// 发送HTTP POST请求
		resp, err := http.Post("http://www.baidu.com", "application/json", reader)
		if err != nil {
			log.Printf("请求第三方服务失败: %v", err)
			return
		}
		defer resp.Body.Close()

		// 打印响应状态
		fmt.Println("响应状态:", resp.Status)

	}
	if err := crontab.AddByFunc("helloWorld", "* * * * *", taskFunc); err != nil {
		logger.Debug("error to add crontab task: ", err)
		return
	}
	crontab.Start()
	select {}
}
