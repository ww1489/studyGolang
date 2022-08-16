package main

import (
	"github.ww.com/Day05/wlog"
	"time"
)

func main() {
	//log := wlog.NewLog("error")
	log := wlog.NewFileLogger("INFO", "./", "wakanda.log", 100*1024*1024)
	msg := "connect Time out"
	defer log.FileClose()
	for {
		log.Debug("这是一条 Debug 日志")
		time.Sleep(time.Second * 1)
		log.Info("这是一条 Info 日志")
		time.Sleep(time.Second * 1)
		log.Worning("这是一条 Worning 日志")
		time.Sleep(time.Second * 1)
		log.Error("MySQL 连接错! Err: %s", msg)
		time.Sleep(time.Second * 1)
		log.Fatal("这是一条 Fatal 日志")
		time.Sleep(time.Second * 1)
	}

}
