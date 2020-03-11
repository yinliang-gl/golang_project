//JSON解析到结构体
package main

import (
	"encoding/json"
	"fmt"
)

type UsergrowthDhhDeliveryAskResponse struct {
	Result  bool
	Task_id string
	Errcode int
}

type ErrorResponse struct {
	Sub_msg  string
	Code     int
	Sub_code string
	Msg      string
}

type Serverslice struct {
	Item                                 string
	Usergrowth_dhh_delivery_ask_response UsergrowthDhhDeliveryAskResponse
	Error_response                       ErrorResponse
}

func main() {
	var s Serverslice
	str1 := `{"item":"fasfds","usergrowth_dhh_delivery_ask_response":
   {"result":true,"task_id":"123456","errcode":1},"error_response":
   {"sub_msg":"非法参数","code":50, "sub_code":"isv.invalid-parameter",  "msg":"Remote service error" }}`

	err := json.Unmarshal([]byte(str1), &s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	// 	fmt.Println(s.Servers[0].ServerName)
}
