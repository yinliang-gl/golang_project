package test_go

import (
	"fmt"
	"log"
	"os"
	"testing"
)

// 简单输出
func TestLog01(t *testing.T) {
	arr := []int{2, 3}
	log.Print("Print array ", arr, "\n")
	log.Println("Println array", arr)
	log.Printf("Printf array with item [%d,%d]\n", arr[0], arr[1])
}

/**
测试 Fatal
对于 log.Fatal 接口，会先将日志内容打印到标准输出，接着调用系统的 os.exit(1) 接口，退出程序并返回状态 1 。
但是有一点需要注意，由于是直接调用系统接口退出，defer函数不会被调用
*/
func TestLog02(t *testing.T) {
	defer func() {
		fmt.Println("--first--")
	}()
	log.Fatalln("test for defer Fatal")
}

/**
测试 Panic
可以看到首先输出了“test for defer Panic”，然后第一个defer函数被调用了并输出了“--first--”，
但是第二个defer 函数并没有输出，可见在Panic之后声明的defer是不会执行的。
*/
func TestLog03(t *testing.T) {

	defer func() {
		fmt.Println("--first--")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	log.Panicln("test for defer Panic")
	defer func() {
		fmt.Println("--second--")
	}()

}

/**
日志写入文件
*/
func TestLog04(t *testing.T) {
	fileName := "test_project/test_go/Info_First.log"
	logFile, err := os.Create(fileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error")
	}
	debugLog := log.New(logFile, "[Info]", log.Llongfile)
	debugLog.Println("A Info message here")
	debugLog.SetPrefix("[Debug]")
	debugLog.Println("A Debug Message here ")
}
