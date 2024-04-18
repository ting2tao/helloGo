package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"strconv"
	"time"
)

var Infof = logrus.Infof

// Logger 中间件函数，用于记录请求信息
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 设置example变量
		c.Set("example", "12345")

		// 请求前
		fmt.Println("开始处理请求")

		c.Next() // 处理请求

		// 请求后
		latency := time.Since(t)
		fmt.Println("请求处理完成")
		status := c.Writer.Status()
		fmt.Printf("状态码: %d, 耗时: %v\n", status, latency)
	}
}

func Logrus(serverName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		setTraceID(c, serverName)
		SetUserID(c)
		startTime := time.Now()
		rawData, _ := c.GetRawData()
		c.Request.Body = io.NopCloser(bytes.NewBuffer(rawData))

		c.Next()
		//endTime := time.Now()
		latencyTime := time.Since(startTime)
		reqMethod := c.Request.Method

		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		Infof("|%3d|%13v|%15s|%s|%s|%s|trace_id=%s",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
			string(rawData),
			c.MustGet("trace_id"),
		)
	}
}

func setTraceID(c *gin.Context, serverName string) {
	traceID := c.GetHeader("trace_id")
	c.Set("server_name", serverName)
	if traceID == "" {
		c.Set("trace_source", serverName)
		traceID = uuid.New().String()
	}
	c.Set("trace_id", traceID)
}

func SetUserID(c *gin.Context) {
	sUserID := c.GetHeader("userID")
	userID, err := strconv.Atoi(sUserID)
	if err != nil {
		log.Print(fmt.Sprintf("异常的用户id:%v", sUserID))
	}
	c.Set("userID", userID)
}

func main() {
	r := gin.Default()
	r.Use(Logger())      // 使用自定义的Logger中间件
	r.Use(Logrus("asd")) // 使用自定义的Logger中间件

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		fmt.Println(example)
	})
	r.POST("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		fmt.Println(example)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
