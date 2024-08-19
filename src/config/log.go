package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"path/filepath"
	"time"
)

var (
	Log          *logrus.Logger
	LogAllConfig = &LogConfig{}
)

// LoggerMiddleware 是一个 Gin 中间件，用于使用 Logrus 记录 HTTP 请求日志
func LoggerMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		//
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		logInfoStr := fmt.Sprintf("%d----%s----[%s]:%s   latencyTime:%v", statusCode, reqMethod, clientIP, reqUri, latencyTime)
		logger.Info(logInfoStr)
	}
}

// 初始化日志logrus
func initLogConfig() {
	LogAllConfig = &ProjectConfig.Logs

	// 配置logrus
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// 设置日志级别
	log.SetLevel(getLogLevel(LogAllConfig.Level))

	// 创建不同级别的日志文件
	infoLog := newLumberjackLogger("info")
	warnLog := newLumberjackLogger("warn")
	errorLog := newLumberjackLogger("error")
	panicLog := newLumberjackLogger("panic")

	// 创建多输出 writer  输出屏幕和输出到文件  os.Stdout
	infoWriter := io.MultiWriter(infoLog)
	warnWriter := io.MultiWriter(warnLog)
	errorWriter := io.MultiWriter(errorLog)
	panicWriter := io.MultiWriter(panicLog)

	// 添加不同级别的 hook   --> 鸭子模型
	log.AddHook(newLogHook(logrus.InfoLevel, infoWriter))
	log.AddHook(newLogHook(logrus.WarnLevel, warnWriter))
	log.AddHook(newLogHook(logrus.ErrorLevel, errorWriter))
	log.AddHook(newLogHook(logrus.PanicLevel, panicWriter))

	Log = log
}

type logHook struct {
	levels []logrus.Level
	writer io.Writer
}

func newLogHook(level logrus.Level, writer io.Writer) *logHook {
	return &logHook{
		levels: []logrus.Level{level},
		writer: writer,
	}
}

func (hook *logHook) Levels() []logrus.Level {
	return hook.levels
}

func (hook *logHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	_, err = hook.writer.Write([]byte(line))
	return err
}

func newLumberjackLogger(level string) *lumberjack.Logger {
	today := time.Now().Format(LogAllConfig.DayFormat)
	logFolder := filepath.Join(LogAllConfig.Dir, today)
	logFile := filepath.Join(logFolder, fmt.Sprintf("%s.log", level))

	// Logger is an io.WriteCloser that writes to the specified filename.
	return &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	}
}

func getLogLevel(level string) logrus.Level {
	switch level {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	case "panic":
		return logrus.PanicLevel
	default:
		return logrus.TraceLevel
	}
}
