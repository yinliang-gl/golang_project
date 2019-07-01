package main

import "os"

/**
从stderr 输出错误信息
*/
func print_test_1() {
	os.Stderr.WriteString("Message")
}
func main() {
	print_test_1()
}
