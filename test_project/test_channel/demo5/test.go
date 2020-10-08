package main

import (
	"fmt"
)

/*
测试channel用于通知中断退出的问题
*/
func testQuit() {
	g := make(chan int)
	quit := make(chan bool)
	go func() {
		for {
			select {
			case v := <-g:
				fmt.Println(v)
			case <-quit:
				fmt.Println("B退出")
				return
			}
		}
	}()
	for i := 0; i < 3; i++ {
		g <- i
	}
	quit <- true
	fmt.Println("testAB退出")
}

func main() {
	testQuit()
}
