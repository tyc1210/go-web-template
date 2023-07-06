package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"go-web-template/internal/setting"
	"io"
	"os"
	"time"
)

var log *logrus.Logger

func NewLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		// 执行时间
		latency := end.Sub(start)
		// 请求路径
		path := c.Request.URL.Path
		// 请求方法
		method := c.Request.Method
		// 状态码
		statusCode := c.Writer.Status()
		userId, _ := c.Get("userId")
		// 日志消息
		log.WithFields(logrus.Fields{
			"userId":     userId,
			"time":       end.Format("2006-01-02 15:04:05"), // 请求结束时间
			"usedTime":   latency.String(),                  // 执行时间
			"path":       path,                              // 请求路径
			"method":     method,                            // 请求方法
			"statusCode": statusCode,                        // 状态码
		}).Info("gin request") // 记录日志
	}
}

func InitLogger(config *setting.AppProperties) {
	log = logrus.New()
	// 设置日志级别
	level, _ := logrus.ParseLevel(config.LoggerLevel)
	log.SetLevel(level)
	//设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	log.SetFormatter(&logrus.TextFormatter{
		ForceQuote:      true,                  //键值对加引号
		TimestampFormat: "2006-01-02 15:04:05", //时间格式
		FullTimestamp:   true,
	})
	// 使用 lumberjack 进行日志轮换
	logger := &lumberjack.Logger{
		Filename:   config.LoggerFileName,   // 日志文件名
		MaxSize:    config.LoggerMaxSize,    // 每个日志文件最大尺寸，单位 MB
		MaxBackups: config.LoggerMaxBackups, // 最多保留的旧日志文件数
		MaxAge:     config.LoggerMaxAge,     // 保留最近的几天日志文件
		LocalTime:  true,                    // 使用本地时间
		Compress:   true,                    // 压缩旧日志文件
	}
	// 设置文件和控制台都进行输出
	log.SetOutput(io.MultiWriter(logger, os.Stdout))
}

// Debug 记录一条 Debug 级别的日志
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// Info 记录一条 Info 级别的日志
func Info(args ...interface{}) {
	log.Info(args...)
}

// Warning 记录一条 Warning 级别的日志
func Warning(args ...interface{}) {
	log.Warning(args...)
}

// Error 记录一条 Error 级别的日志
func Error(args ...interface{}) {
	log.Error(args...)
}
