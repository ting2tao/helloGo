package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"hello/pkg/cron"
	"hello/pkg/logger"
	"io"
	"log"
	"net/http"
)

// 自定义的 ResponseWriter，用于捕获响应数据
type CustomResponseWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

// Write 方法的重写，用于捕获数据
func (w *CustomResponseWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

// 记录第三方请求和响应的中间件
func ThirdPartyRequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 发起到第三方的请求前
		requestData, _ := c.GetRawData()
		// 重新设置请求体，以便后续处理
		c.Request.Body = io.NopCloser(bytes.NewBuffer(requestData))

		// 记录请求信息
		log.Printf("请求第三方服务: %s, 请求数据: %s", c.Request.URL.String(), string(requestData))
		// 创建一个新的 CustomResponseWriter
		bodyWriter := &CustomResponseWriter{Body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter
		// 请求处理
		c.Next()

		// 发起到第三方的请求后
		responseBody := bodyWriter.Body.String()

		// 记录响应信息
		log.Printf("第三方服务响应: %s", responseBody)
	}
}

func thirdParty() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("thirdParty")
	}
}

func main() {
	r := gin.Default()
	r.Use(ThirdPartyRequestLogger()) // 使用中间件

	// 定义路由
	r.POST("/example", thirdParty())

	r.POST("/example/v1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})
	go Task()
	// 启动服务器
	r.Run(":8080")

}

func Task() {
	crontab := cron.NewCrontab()
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
	// 添加函数作为定时任务
	taskFunc := func() {
		// 代理到第三方服务
		resp, err := http.Post("https://www.baidu.com", "application/json", reader)
		if err != nil {
			logger.Debug("请求第三方服务失败: %v", err)
			//c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		fmt.Println(resp)
	}
	if err := crontab.AddByFunc("helloWorld", "* * * * *", taskFunc); err != nil {
		logger.Debug("error to add crontab task: ", err)
		return
	}
	crontab.Start()
	select {}
}
