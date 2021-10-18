package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"io/fs"
	logsys "log"
	"os"
	"path"
	"time"
)

// LogError 当存在错误时记录日志
func LogError(err error) {
	if err != nil {
		logsys.Println(err)
	}
}

func Logger(logPath string) gin.HandlerFunc {
	logClient := log.New()
	//var logPath = "/var/log/katy" // 日志打印到指定的目录
	//var logPath = config.GetString("app.log") // 日志打印到指定的目录
	createPath(logPath, os.ModePerm)
	// 目录不存在则创建

	fileName := path.Join(logPath, "gin-api.log")
	//禁止logrus的输出
	//src, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	LogError(err)
	// 设置日志输出的路径
	logClient.Out = src
	logClient.SetLevel(log.DebugLevel)
	//apiLogPath := "gin-api.log"
	logWriter, _ := rotatelogs.New(
		fileName+".%Y-%m-%d-%H-%M-%S.log",
		rotatelogs.WithLinkName(fileName),         // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(365*24*time.Hour),     // 文件最大保存时间
		//rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
		rotatelogs.WithRotationTime(time.Second), // 日志切割时间间隔
	)
	writeMap := lfshook.WriterMap{
		log.InfoLevel:  logWriter,
		log.FatalLevel: logWriter,
		log.DebugLevel: logWriter, // 为不同级别设置不同的输出目的
		log.WarnLevel:  logWriter,
		log.ErrorLevel: logWriter,
		log.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &log.JSONFormatter{})
	logClient.AddHook(lfHook)

	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)

		path := c.Request.URL.Path

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		// 这里是指定日志打印出来的格式。分别是状态码，执行时间,请求ip,请求方法,请求路由(等下我会截图)
		logClient.Infof("| %3d | %13v | %15s | %s  %s |",
			statusCode,
			latency,
			clientIP,
			method, path,
		)
	}
}

func createPath(uPath string, pem fs.FileMode) bool {
	if isExist(uPath) == false {
		err := os.MkdirAll(uPath, pem)
		if err != nil {
			LogError(err)
			return false
		}
	}
	return true
}

//判断文件夹是否存在
func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			LogError(err)
			return false
		}
		//fmt.Println(err)
		return false
	}
	return true
}
