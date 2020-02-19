package test_context

// 参考： https://blog.csdn.net/u011957758/article/details/82948750

import (
	"fmt"
	"golang.org/x/net/context"
	"testing"
	"time"
)

/**
吃汉堡比赛，奥特曼每秒吃0-5个，计算吃到10的用时
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

/**
吃汉堡比赛，奥特曼每秒吃0-5个，用时10秒，可以吃多少个
*/
func TestContext002(t *testing.T) {
	// ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	chiHanBao_02(ctx)
	defer cancel()
}

/**
携带关键信息，为全链路提供线索，比如接入elk等系统，需要来一个trace_id，那WithValue就非常适合做这个事。
*/
func TestContext003(t *testing.T) {
	ctx := context.WithValue(context.Background(), "trace_id", "88888888")
	// 携带session到后面的程序中去
	ctx = context.WithValue(ctx, "session", 1)

	process(ctx)

}
