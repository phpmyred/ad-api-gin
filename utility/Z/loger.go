package Z

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var LoggerObj *logrus.Logger

// LogFormatter 日志自定义格式
type logFormatter struct{}

// Format
// @Description 格式详情设置日志的格式
// @Author aDuo 2024-08-15 17:41:23
// @Param entry
// @Return []byte
// @Return error
func (s *logFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006-01-02 15:04:05")
	var file string
	var l int
	if entry.Caller != nil {
		file = filepath.Base(entry.Caller.File)
		l = entry.Caller.Line
	}
	msg := fmt.Sprintf("%s [%s:%d][GOID:%d][%s] %s\n", timestamp, file, l, getGID(), strings.ToUpper(entry.Level.String()), entry.Message)

	return []byte(msg), nil
}

// getGID
// @Description 获取GO的进程
// @Author aDuo 2024-08-15 17:42:04
// @Return uint64
func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

// init
// @Description 初始化日志
// @Author aDuo 2024-08-15 17:42:28
func init() {
	if LoggerObj != nil {
		src, _ := setOutputFile()
		//设置输出
		LoggerObj.Out = src
		return
	}
	//实例化
	logger := logrus.New()
	src, _ := setOutputFile()
	//设置输出
	logger.Out = src
	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	/*
		加个hook形成ELK体系
	*/
	//hook := model.EsHookLog()
	//logger.AddHook(hook)
	logger.SetReportCaller(true)
	logger.SetFormatter(new(logFormatter))

	LoggerObj = logger
}

// setOutputFile
// @Description  设置日志的存放路径
// @Author aDuo 2024-08-15 17:42:40
// @Return *os.File
// @Return error
func setOutputFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(logFilePath, 0777); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	logFileName := now.Format("2006-01-02") + ".log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return src, nil
}
