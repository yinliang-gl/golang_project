package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	//当前时间戳
	fmt.Println(time.Now().Format(time.RFC3339Nano), time.Now().Unix())
	format_string := "2006-01-02T15:04:05"
	fmt.Println(time.Now().Format(format_string), time.Now().Unix())

	datetime := time.Now().Format("2006-01-02T15:04:05") //后面的参数是固定的 否则将无法正常输出
	str := datetime + "+08:00"
	fmt.Println(str)

	log.Printf("\n{\"langyan\":\"监控\",\"topic\":\"alarm_log01\",\"messages\":[{\"name\":\"databus_complete\",\"tags\":{\"task\":\"index_dumper_benchmark\"},\"values\":{\"c\":0,\"cost\":%d},\"time\":\"%s\"}]}", 123, "aaaaaaaaaaaaaaaaaa")
}
