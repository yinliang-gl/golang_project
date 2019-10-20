package test_go

import (
	"log"
	"os"
	"testing"
)

/**
从stderr 输出错误信息
*/
func TestPrint(t *testing.T) {
	_, _ = os.Stderr.WriteString("Message\n")

	l := log.New(os.Stderr, "fdasfdsf:", 0)
	l.Println("aaaaaa")
}
