package main

// 定制的日志记录器

// 这个示例程序展示如何创建定制的日志记录器

import (
	"fmt"
	"git.qutoutiao.net/CPC/alarm-sdk/jlog"
	"log"
	"os"
	"time"
)

var (
	Error *log.Logger // 非常严重的问题
)

func main() {
	name, err1 := os.Hostname()
	if err1 != nil {
		log.Fatalf("get host name failed")
	}
	err := jlog.Init(name, "indexDumpError", "index_dump.ERROR", []string{"warn", "trace", "error"}, map[string]string{"default": "alarmtest"})
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		jlog.JError1("error", "fasfds")
		time.Sleep(1 * time.Second)
	}
}
