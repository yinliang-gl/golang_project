package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取主机名
	if name, err := os.Hostname(); err == nil {
		fmt.Println("获取主机名成功，主机名为 " + name)
	} else {
		fmt.Println("获取主机名失败")
	}

}
