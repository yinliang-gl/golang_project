package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

// 标准的 client （充分使用复用技术）
// 参考   https://studygolang.com/articles/17716
func main() {
	url := `http://httpbin.org/post?key=123`

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req) // 用完需要释放资源

	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	req.SetRequestURI(url)

	requestBody := []byte(`{"request":"test"}`)
	req.SetBody(requestBody)

	req.SetConnectionClose()

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp) // 用完需要释放资源

	if err := fasthttp.DoTimeout(req, resp, 5*time.Second); err != nil {
		fmt.Println("请求失败:", err.Error())
		return
	}

	b := resp.Body()

	fmt.Println("result:\r\n", string(b))
}
