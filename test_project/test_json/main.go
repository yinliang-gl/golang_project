//JSON解析到结构体
package main

/*
	这种json解析速度没有json-iterator快，json-iterator参考test_json-iterator
*/
import (
	"encoding/json"
	"fmt"
)

type OperatorItem struct {
	Key       uint64  `json:"key"`
	Op        string  `json:"op"`
	Value     []int64 `json:"value"`
	Threshold string  `json:"threshold"`
}

type OperatorList struct {
	operatorItems []OperatorItem `json:"list"`
}

func main() {
	var err error







	str1 := `{
      "key": "budget",
      "op": "inc",
      "value": [10],
      "thres": 0
    }`

	var OperatorItem OperatorItem
	err = json.Unmarshal([]byte(str1), &OperatorItem)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", OperatorItem)

	return

}

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

type TaobaoRtaResponse struct {
	Usergrowth_dhh_delivery_ask_response UsergrowthDhhDeliveryAskResponse
	Error_response                       ErrorResponse
}

func main1() {
	var s TaobaoRtaResponse
	//str := `{"usergrowth_dhh_delivery_ask_response":
	//{"result":true,"task_id":"123456","errcode":1},"error_response":
	//{"sub_msg":"非法参数","code":50, "sub_code":"isv.invalid-parameter",  "msg":"Remote service error" }}`

	str := `{"usergrowth_dhh_delivery_ask_response":
	{"result":true,"task_id":"123456","errcode":1}}`

	//str := `{"error_response":
	//{"sub_msg":"非法参数","code":50, "sub_code":"isv.invalid-parameter",  "msg":"Remote service error" }}`

	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	// 	fmt.Println(s.Servers[0].ServerName)
}

type RuleItem struct {
	Level       string  `json:"level"`
	DayInterval int     `json:"day_interval"`
	Key         string  `json:"key"`
	Op          string  `json:"op"`
	Value       []int64 `json:"value"`
}

type RuleAnd struct {
	Rules []RuleItem `json:"rules"`
}

type RuleAndOr struct {
	RuleAnds []RuleAnd `json:"list"`
}

func main2() {
	var str string
	var err error

	var s RuleAndOr

	str = `{
  "list": [
    {
      "rules": [
        {
          "level": "user",
          "day_interval": 1,
          "key": "inner_realtime_active",
          "op": "gt",
          "value": [
            100
          ]
        },
        {
          "level": "group",
          "day_interval": 1,
          "key": "inner_realtime_active",
          "op": "gt",
          "value": [
            10
          ]
        }
      ]
    },
    {
      "rules": [
        {
          "level": "user",
          "day_interval": 1,
          "key": "inner_realtime_active",
          "op": "gt",
          "value": [
            100
          ]
        }
      ]
    }
  ]
}`
	err = json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Printf("%#v", s)
	}
	// fmt.Println(s)
	fmt.Printf("%#v", s)

	return

	str = `{
    "rules": [
      {
        "level": "user",
        "day_interval": 1,
        "key": "inner_realtime_active",
        "op": "gt",
        "value": [
          100
        ]
      },
      {
        "level": "group",
        "day_interval": 1,
        "key": "inner_realtime_active",
        "op": "gt",
        "value": [
          10
        ]
      }
    ]
  }`

	var ruleAnd RuleAnd
	err = json.Unmarshal([]byte(str), &ruleAnd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", ruleAnd)
	fmt.Printf("%#v\n", len(ruleAnd.Rules))

	return

	str = `{"level":"user","day_interval":1,"key":"inner_realtime_active","op":"gt","value":[100]}`
	var item RuleItem
	err = json.Unmarshal([]byte(str), &item)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", item)
	return

}
