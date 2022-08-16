package wlog

import (
	"fmt"
	"os"
	"path"
)

// 定义文件存储日志结构体
type fileLogger struct {
	level      LogLevel
	Filepath   string
	Filename   string
	fileobj    *os.File
	errfileobj *os.File
	Filesize   int64
}

// NewFileLogger 定义一个文件存储日志结构体构建函数
func NewFileLogger(le, fp, fn string, fz int64) *fileLogger {
	level, err := parseLogLevel(le)
	if err != nil {
		panic(err)
	}
	f1 := &fileLogger{
		level:    level,
		Filepath: fp,
		Filename: fn,
		Filesize: fz,
	}
	err = f1.initFile()
	if err != nil {
		panic(err)
	}
	return f1
}

func (f *fileLogger) initFile() error {
	fullFileNme := path.Join(f.Filepath, f.Filename)
	fileObj, err := os.OpenFile(fullFileNme+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Read file from failed", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileNme+"err"+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Read file from failed", err)
		return err
	}
	f.fileobj = fileObj
	f.errfileobj = errFileObj
	return nil
}

// message 格式化输出的消息
func (f *fileLogger) message(leves, format string, a ...interface{}) {
	lv, err := parseLogLevel(leves)
	if err != nil {
		return
	}
	g := fmt.Sprintf(format, a...)
	funcname, filename, lines := getInfo(3)
	fmt.Fprintf(f.fileobj, " [%s] [%s:%s:%d] [%s] msg: %s\n", now, filename, funcname, lines, leves, g)
	if lv >= ERROR {
		fmt.Fprintf(f.errfileobj, " [%s] [%s:%s:%d] [%s] msg: %s\n", now, filename, funcname, lines, leves, g)
	}
}

func (f *fileLogger) Debug(format string, a ...interface{}) {
	if f.level <= DEBUG {
		f.message("DEBUG", format, a...)
	}
}

func (f *fileLogger) Info(format string, a ...interface{}) {
	if f.level <= INFO {
		f.message("INFO", format, a...)
	}
}

func (f *fileLogger) Worning(format string, a ...interface{}) {
	if f.level <= WORNING {
		f.message("WORNING", format, a...)
	}
}

func (f *fileLogger) Error(format string, a ...interface{}) {
	if f.level <= ERROR {
		f.message("ERROR", format, a...)
	}
}

func (f *fileLogger) Fatal(format string, a ...interface{}) {
	if f.level <= FATAL {
		f.message("FATAL", format, a...)
	}
}

func (f *fileLogger) FileClose() {
	f.fileobj.Close()
	f.errfileobj.Close()
}
