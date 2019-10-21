package main

import (
	"fmt"
	"time"
)

/*
测试通过channel来控制最大并发数，来处理事件
*/
func testMaxNumControl() {
	maxNum := 3
	limit := make(chan bool, maxNum)
	quit := make(chan bool)

	for i := 0; i < 100; i++ {
		fmt.Println("start worker : ", i)

		limit <- true

		go func(i int) {
			fmt.Println("do worker start: ", i)
			time.Sleep(time.Millisecond * 20)
			fmt.Println("do worker finish: ", i)
			<-limit
			if i == 99 {
				fmt.Println("完成任务")
				quit <- true
			}
		}(i)
	}

	<-quit
	fmt.Println("收到退出通知，主程序退出")
}
func main() {
	testMaxNumControl()
}
