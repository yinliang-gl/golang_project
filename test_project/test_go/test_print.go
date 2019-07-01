package main

import (
	"log"
	"os"
)

/**
从stderr 输出错误信息
*/
func print_test_1() {
	_, _ = os.Stderr.WriteString("Message\n")

	l := log.New(os.Stderr, "fdasfdsf:", 0)
	l.Println("aaaaaa")
}
func main() {
	print_test_1()
}
