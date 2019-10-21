package main

import (
	"fmt"
	"time"
)

/*
指定channel是输入还是输出型的，防止编写时写错误输入输出，指定了的话，可以在编译时期作错误的检查
*/
func testInAndOutChan() {
	ch := make(chan int)
	quit := make(chan bool)

	//输入型的chan是这种格式的：inChan chan<- int，如果换成输出型的，则编译时会报错
	go func(inChan chan<- int) {
		for i := 0; i < 10; i++ {
			inChan <- i
			time.Sleep(time.Millisecond * 500)
		}
		quit <- true
		quit <- true
	}(ch)

	go func(outChan <-chan int) {
		for {
			select {
			case v := <-outChan:
				fmt.Println("print out value : ", v)
			case <-quit:
				fmt.Println("收到退出通知，退出")
				return
			}
		}
	}(ch)

	<-quit
	fmt.Println("收到退出通知，主线程退出")
}
func main() {
	testInAndOutChan()
}
