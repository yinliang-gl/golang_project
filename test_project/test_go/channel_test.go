package test_go

// TODO
import (
	"fmt"
	"testing"
	"time"
)

/**
go关键字可以用来开启一个goroutine(协程))进行任务处理，而多个任务之间如果需要通信，就需要用到channel了。
*/
func TestChannel01(t *testing.T) {
	intChan := make(chan int)

	go func() {
		time.Sleep(10 * time.Second)
		intChan <- 1
	}()

	value := <-intChan
	fmt.Println("value : ", value)
}
