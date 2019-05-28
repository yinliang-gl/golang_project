package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 获取CPU核心数
	fmt.Println(runtime.NumCPU())

	//设置cpu个线程去运行所有的协程
	runtime.GOMAXPROCS(runtime.NumCPU())

}
