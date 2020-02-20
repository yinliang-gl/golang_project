package test_hystrix

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"net/http"
	"testing"
	"time"
)

/*
	通用的调用方式
*/
func TestHystrix001(t *testing.T) {
	hystrix.ConfigureCommand("my_command", hystrix.CommandConfig{
		Timeout:                500,   // 500ms超时
		MaxConcurrentRequests:  10000, // 不限制并发数量
		RequestVolumeThreshold: 3000,  // 10s 内到 3000 的失败就熔断
		SleepWindow:            3000,  // ms 熔断3S
		ErrorPercentThreshold:  25,    // 25% 的请求失败
	})

	output := make(chan bool, 1)
	errors := hystrix.Go("my_command", func() error {
		// talk to other services
		time.Sleep(15 * time.Second)
		output <- true
		return nil
	}, func(err error) error {
		fmt.Println("get an error, handle it, error is ", err)
		return err
	})

	select {
	case out := <-output:
		// success
		fmt.Println("succcess: ", out)
	case err := <-errors:
		// failure
		fmt.Println("failure: ", err)
	}
}

/**
发http请求
*/

func TestHystrix002(t *testing.T) {
	hystrix.Go("get_baidu", func() error {
		// talk to other services
		res, err := http.Get("https://www.baidu.com/")
		if err != nil {
			fmt.Println("get error")
			return err
		}
		//p := []byte{1, 2, 3, 4, 5, 6}
		p := make([]byte, 5000)
		res.Body.Read(p)
		fmt.Println(string(p))
		return nil
	}, func(err error) error {
		fmt.Println("get an error, handle it")
		return nil
	})

	time.Sleep(2 * time.Second) // 调用Go方法就是起了一个goroutine，这里要sleep一下，不然看不到效果
}
