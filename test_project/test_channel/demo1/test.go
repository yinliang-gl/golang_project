package main

import (
	"fmt"
	"time"
)

func testSimple() {
	intChan := make(chan int)

	go func() {
		time.Sleep(10 * time.Second)
		intChan <- 2
	}()

	value := <-intChan
	fmt.Println("value : ", value)
}

func main() {
	testSimple()
}
