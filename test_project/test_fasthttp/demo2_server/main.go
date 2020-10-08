package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
)

// 参考   https://kuaibao.qq.com/s/20191029A0A4RE00?refer=cp_1026
func main() {

	log.Fatal(fasthttp.ListenAndServe(":8001", func(ctx *fasthttp.RequestCtx) {
		path := string(ctx.Path())
		switch path {
		case "/post":
			httpHandle(ctx)
		default:
			fmt.Println("---------------------")
		}
	}))
}

func httpHandle(ctx *fasthttp.RequestCtx) {
	// 定义新的request
	newRequest := &fasthttp.Request{}
	// 复制
	ctx.Request.CopyTo(newRequest)

	// 获取请求body
	body := newRequest.Body()
	fmt.Println(string(body))

	newRequest = nil
	ctx.Response.AppendBodyString("ok")
	ctx.Response.SetStatusCode(204)

}
