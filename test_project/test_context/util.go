package test_context

import (
	"fmt"
	"golang.org/x/net/context"
	"math/rand"
	"time"
)

func chiHanBao_01(ctx context.Context) <-chan int {
	c := make(chan int)
	// 个数
	n := 0
	// 时间
	t := 0
	go func() {
		for {
			//time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				fmt.Printf("耗时 %d 秒，吃了 %d 个汉堡 \n", t, n)
				return
			case c <- n:
				incr := rand.Intn(5)
				n += incr
				t++
				fmt.Printf("我吃了 %d 个汉堡\n", n)
			}
		}
	}()

	return c

}

func chiHanBao_02(ctx context.Context) {
	n := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop")
			return
		default:
			incr := rand.Intn(5)
			n += incr
			fmt.Printf("我吃了 %d 个汉堡\n", n)
		}
		time.Sleep(time.Second)
	}
}

func process(ctx context.Context) {
	session, ok := ctx.Value("session").(int)
	fmt.Println(ok)
	if !ok {
		fmt.Println("something wrong")
		return
	}

	if session != 1 {
		fmt.Println("session 未通过")
		return
	}

	traceID := ctx.Value("trace_id").(string)
	fmt.Println("trace_id:", traceID, ", session:", session)
}
