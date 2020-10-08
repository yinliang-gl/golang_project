package main

import (
	"fmt"
	"time"
)

/**
生产者消费者问题
*/
func testPCB() {
	fmt.Println("test PCB")

	intchan := make(chan int)
	quitChan := make(chan bool)
	quitChan2 := make(chan bool)

	value := 0
	go func() {
		for i := 0; i < 3; i++ {
			value = value + 1
			intchan <- value
			fmt.Println("write finish, value ", value)
			time.Sleep(time.Second)
		}
		quitChan <- true
	}()
	go func() {
		for {
			select {
			case v := <-intchan:
				fmt.Println("read finish, value ", v)
			case <-quitChan:
				quitChan2 <- true
				return
			}
		}

	}()
	<-quitChan2
	fmt.Println("task is done ")
}

func main() {
	testPCB()
}
