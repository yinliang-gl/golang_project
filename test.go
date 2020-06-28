package main

import (
	"fmt"
	"time"
)

func main() {
	str := []string{"1", "2", "3"}
	fmt.Println(str[:0])

	fmt.Println(time.Now().Unix())
}
