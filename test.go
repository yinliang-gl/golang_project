package main

import (
	"fmt"
	"time"
)

func main() {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(timeStr)

	timeStr = string([]byte(timeStr)[11:])
	fmt.Println(timeStr)
}
