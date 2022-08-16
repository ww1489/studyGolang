package wlog

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

// LogLevel 自定义一个类型
type LogLevel uint16

// Logger 日志结构体
type Logger struct {
	level LogLevel
}

const (
	DEBUG LogLevel = iota
	INFO
	WORNING
	ERROR
	FATAL
	UNKNOW
)

// 构建一个接口，方便实现往stdout 和 文件打印
//type outType interface {
//	Debug()
//	Info()
//	Worning()
//	Error()
//	Fatal()
//}

var now = time.Now().Format("2006-01-02 15:04:05")

// parseLogLevel 构建日志等级，方便打印
func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "worning":
		return WORNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的等级")
		return UNKNOW, err
	}
}

// getInfo 获取文件信息，实现打印行数功能
func getInfo(skip int) (funcName, fileName string, line int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("Runtime.Caller() Failed!")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return
}

// message 格式化输出的消息
func (l Logger) message(leves, format string, a ...interface{}) {
	g := fmt.Sprintf(format, a...)
	funcname, filename, lines := getInfo(3)
	fmt.Printf(" [%s] [%s:%s:%d] [%s] msg: %s\n", now, filename, funcname, lines, leves, g)
}
