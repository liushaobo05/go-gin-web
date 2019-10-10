package middleware

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}

// 日志输出到控制台
func LoggerToConsole() gin.HandlerFunc {
	//实例化
	logger := logrus.New()

	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	//设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	//设置输出
	logger.SetOutput(os.Stdout)

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求id
		requestId := GetReqID(c)

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		logger.WithFields(logrus.Fields{
			"statusCode":  statusCode,
			"requestId":   requestId,
			"latencyTime": latencyTime,
			"clientIP":    clientIP,
			"reqMethod":   reqMethod,
			"reqUri":      reqUri,
		}).Info()
	}
}

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	//实例化
	loggerConsole := logrus.New()

	//设置日志级别
	loggerConsole.SetLevel(logrus.DebugLevel)

	//设置日志格式
	loggerConsole.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	//设置输出
	loggerConsole.SetOutput(os.Stdout)

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求id
		requestId := GetReqID(c)

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		loggerConsole.Infof("code = %3d, reqId = %s, reqCost = %v, reqIp = %s, method = %s, url = %s",
			statusCode,
			requestId,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

// 日志记录到 MongoDB
func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
