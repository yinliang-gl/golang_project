package main

import (
	"fmt"
	"time"
)

func main() {
	//当前时间戳
	t1 := time.Now().Unix()
	fmt.Println(t1)

	time.Sleep(10 * time.Second)
	t2 := time.Now().Unix()
	fmt.Println(t2)

	fmt.Print(t2 - t1)
}
