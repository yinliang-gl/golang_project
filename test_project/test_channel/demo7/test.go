package main

import (
	"fmt"
	"time"
)

/*
检查channel读写超时，并做超时的处理
*/
func testTimeout() {
	g := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case v := <-g:
				fmt.Println(v)
			case <-time.After(time.Second * time.Duration(3)):
				quit <- true
				fmt.Println("超时，通知主线程退出")
				return
			}
		}
	}()

	for i := 0; i < 5; i++ {
		g <- i
	}

	<-quit
	fmt.Println("收到退出通知，主线程退出")
}
func main() {
	testTimeout()
}
