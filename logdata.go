package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"time"
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

type LogData struct {
	RequestID         string          `json:"request_id"`
	RequestURI        string          `json:"request_uri"`
	RequestMethod     string          `json:"request_method"`
	ResponseTime      string          `json:"response_time"`
	StatusCode        int             `json:"status_code"`
	ClientIP          string          `json:"client_ip"`
	Header            string          `json:"header"`
	HeaderBytes       []byte          `json:"header_bytes"`
	RequestBodyBytes  []byte          `json:"request_body_bytes"`
	ResponseBodyBytes []byte          `json:"response_body_bytes"`
	RequestBody       json.RawMessage `json:"request_body"`
	ResponseBody      json.RawMessage `json:"response_body"`
}

// 中间件，记录请求和响应数据，并生成唯一请求ID
func RequestResponseLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()
		// 生成唯一请求ID
		requestID := uuid.New().String()
		c.Set("requestID", requestID)

		rawData, _ := c.GetRawData()
		c.Request.Body = io.NopCloser(bytes.NewBuffer(rawData))

		// 创建一个新的 CustomResponseWriter
		bodyWriter := &CustomResponseWriter{Body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter
		// 创建一个bytes.Buffer来存储headers
		buf := new(bytes.Buffer)

		// 将headers写入buffer中
		c.Request.Header.WriteSubset(buf, nil)

		// 将buffer转换为[]byte
		headersAsBytes := buf.Bytes()
		// 请求处理
		c.Next()

		// 获取响应状态码和响应体
		statusCode := bodyWriter.Status()
		responseBody := bodyWriter.Body.String()
		if len(rawData) == 0 {
			rawData = json.RawMessage(`""`)
		}
		logData := &LogData{
			RequestID:         requestID,
			RequestURI:        c.Request.RequestURI,
			RequestMethod:     c.Request.Method,
			ResponseTime:      time.Since(now).String(),
			StatusCode:        statusCode,
			RequestBody:       rawData,
			ClientIP:          c.ClientIP(),
			ResponseBody:      json.RawMessage(responseBody),
			RequestBodyBytes:  rawData,
			ResponseBodyBytes: []byte(responseBody),
			HeaderBytes:       headersAsBytes,
			Header:            string(headersAsBytes),
		}
		// 记录日志
		log.Printf("ID: %s, 路径: %s,方法: %s,响应时间:%s, 状态码: %d,请求体：%s, 响应体: %s\n", requestID, c.Request.RequestURI, c.Request.Method, time.Since(now), statusCode, rawData, responseBody)
		dataStr, err := json.Marshal(logData)
		if err != nil {
			log.Println("err", err.Error())
		}
		log.Println("logData", string(dataStr))
		// 你可以在这里将日志保存到数据库或文件中
	}
}

func main() {
	r := gin.Default()
	r.Use(RequestResponseLogger())   // 使用中间件
	r.Use(ThirdPartyRequestLogger()) // 使用中间件

	// 定义路由
	r.GET("/example", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})
	r.POST("/example", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	// 启动服务器
	r.Run(":8080")
}

// 记录第三方请求和响应的中间件
func ThirdPartyRequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 发起到第三方的请求前
		requestData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Printf("读取请求体失败: %v", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
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
