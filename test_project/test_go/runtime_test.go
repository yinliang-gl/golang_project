package test_go

import (
	"fmt"
	"runtime"
	"testing"
)

func TestRuntime(t *testing.T) {
	// 获取CPU核心数
	fmt.Println(runtime.NumCPU())

	//设置cpu个线程去运行所有的协程
	runtime.GOMAXPROCS(runtime.NumCPU())

}
