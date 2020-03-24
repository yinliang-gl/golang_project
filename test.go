package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"net"
	"time"
)

func main() {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/x-www-form-urlencoded;charset=utf-8")
	req.SetRequestURI("http://gw.api.taobao.com/router/rest")
	req.PostArgs().Set("app_key", "12129701")
	req.PostArgs().Set("format", "json")
	req.PostArgs().Set("method", "taobao.usergrowth.dhh.delivery.ask")
	req.PostArgs().Set("partner_id", "apidoc")
	req.PostArgs().Set("sign", "2C2D8197838C2B4410DB8F23BF19D7B6")
	req.PostArgs().Set("sign_method", "hmac")
	req.PostArgs().Set("timestamp", "2020-03-10 18:11:31")
	req.PostArgs().Set("v", "2.0")
	req.PostArgs().Set("advertising_space_id", "2187")
	req.PostArgs().Set("channel", "1")
	req.PostArgs().Set("idfa", "76B17F93-F503-4CF4-A05E-7A9FUECDA144")
	req.PostArgs().Set("idfa_md5", "0f7a929446fb7daf16a8eb6f2e160725")
	req.PostArgs().Set("imei", "864162033703542")
	req.PostArgs().Set("imei_md5", "0f7a929446fb7daf16a8eb6f2e160725")
	req.PostArgs().Set("oaid", "77CE527B70CB45AD991D8535BF94388914436004019A9D5BC74C30E80BE184F1")
	req.PostArgs().Set("oaid_md5", "0f7a929446fb7daf16a8eb6f2e160725")
	req.PostArgs().Set("os", "0")
	req.PostArgs().Set("profile", "无")

	fmt.Printf("req:%#v", req)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	Cli := &fasthttp.Client{
		Dial: func(addr string) (conn net.Conn, e error) {
			return fasthttp.DialTimeout(addr, 2000)
		},
	}

	err := Cli.DoTimeout(req, resp, time.Millisecond*2000)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Printf("%d\n", resp.Header.StatusCode())
	fmt.Printf("resp body == %v \n", (string)(resp.Body()))
}
