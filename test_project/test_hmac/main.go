package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
)

func main() {
	test_02()
}
func test_02() {
	argsMap := make(map[string]string)

	argsMap["app_key"] = "12129701"
	argsMap["format"] = "json"
	argsMap["method"] = "taobao.usergrowth.dhh.delivery.ask"
	argsMap["partner_id"] = "apidoc"

	argsMap["sign_method"] = "hmac"
	// argsMap["timestamp"] = "2020-03-10 18:02:04"
	argsMap["v"] = "2.0"
	argsMap["advertising_space_id"] = "2187"
	argsMap["channel"] = "1"
	argsMap["idfa"] = "76B17F93-F503-4CF4-A05E-7A9FUECDA144"
	argsMap["idfa_md5"] = "0f7a929446fb7daf16a8eb6f2e160725"
	argsMap["imei"] = "864162033703542"
	argsMap["imei_md5"] = "0f7a929446fb7daf16a8eb6f2e160725"
	argsMap["oaid"] = "77CE527B70CB45AD991D8535BF94388914436004019A9D5BC74C30E80BE184F1"
	argsMap["oaid_md5"] = "0f7a929446fb7daf16a8eb6f2e160725"
	argsMap["os"] = "0"
	argsMap["profile"] = "无"

	var keys []string
	for k := range argsMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	str := ""
	for _, val := range keys {
		str += val
		str += argsMap[val]
	}
	fmt.Println(str)

	key := []byte(str)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte("aaaaaa"))
	fmt.Printf("%X\n", mac.Sum(nil))
}

func test_01() {
	//sha1
	h := sha1.New()
	io.WriteString(h, "aaaaaa")
	fmt.Printf("%x\n", h.Sum(nil))

	//hmac ,use sha1
	key := []byte("123456")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte("aaaaaa"))
	fmt.Printf("%x\n", mac.Sum(nil))
}
