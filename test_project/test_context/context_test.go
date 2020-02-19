package test_context

// 参考： https://blog.csdn.net/u011957758/article/details/82948750

import (
	"fmt"
	"golang.org/x/net/context"
	"testing"
)

/**

 */
func TestContext001(t *testing.T) {
	ctx, a_cancel := context.WithCancel(context.Background())
	eatNum := chiHanBao_01(ctx)
	for n := range eatNum {
		if n >= 10 {
			a_cancel()
			break
		}
	}

	fmt.Println("正在统计结果。。。")
	//time.Sleep(1 * time.Second)
}
