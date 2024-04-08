package logger

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Logrus(serverName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		setTraceID(c, serverName)
		SetUserID(c)
		startTime := time.Now()
		rawData, _ := c.GetRawData()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rawData))
		deviceID := c.Request.Header.Get("device_id")
		if deviceID == "" {
			deviceID = c.Request.Header.Get("deviceid")
		}
		c.Set("device_id", deviceID)
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		platform := c.Request.Header.Get("platform")
		if platform == "" {
			platform = c.Request.Header.Get("ServerName")
			if platform == "" {
				platform = "Unknown"
			}
		}
		appVersion := c.Request.Header.Get("app_version")
		if appVersion == "" {
			appVersion = c.Request.Header.Get("version")
		}
		osVersion := c.Request.Header.Get("os_version")
		model := c.Request.Header.Get("model")
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		Infof("|%3s|%3s|%3s|%3s|%3s|%3d|%13v|%15s|%s|%s|%s|trace_id=%s",
			model,
			osVersion,
			platform,
			appVersion,
			deviceID,
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
