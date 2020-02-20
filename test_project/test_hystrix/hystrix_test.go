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
		Timeout:                1000 * 10, // 请求超时的时间，单位毫秒
		ErrorPercentThreshold:  1,         // 允许出现的错误比例
		SleepWindow:            10000,     // 熔断开启多久尝试发起一次请求
		MaxConcurrentRequests:  1000,      // 允许的最大并发请求数
		RequestVolumeThreshold: 5,         // 波动期内的最小请求数，默认波动期 10S
	})

	output := make(chan bool, 1)
	errors := hystrix.Go("my_command", func() error {
		// talk to other services
		time.Sleep(15 * time.Second)
		output <- true
		return nil
	}, nil) // 第二个参数不为nil的时候，超时不返回，而是调用第二个参数指定的函数。

	select {
	case out := <-output:
		// success
		fmt.Println("succcess ", out)
	case err := <-errors:
		// failure
		fmt.Println("failure ", err)
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
