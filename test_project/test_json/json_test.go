package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type OperatorItem struct {
	Key       string  `json:"key"`
	Op        string  `json:"op"`
	Value     []int64 `json:"value"`
	Threshold string  `json:"threshold"`
}

type OperatorList struct {
	OperatorItems []OperatorItem `json:"list"`
}

func TestAnalyseJson2(t *testing.T) {
	str := `{
  "list": [
    {
      "key": "budget",
      "op": "inc",
      "value": [
        10
      ],
      "thres": 0
    }
  ]
}`
	// fmt.Println(str)
	var OperatorList OperatorList
	err := json.Unmarshal([]byte(str), &OperatorList)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", OperatorList)
	fmt.Printf("%#v\n", len(OperatorList.OperatorItems))

}

func TestAnalyseJson1(t *testing.T) {
	str1 := `{
      "key": "budget",
      "op": "inc",
      "value": [10],
      "thres": 0
    }`

	var OperatorItem OperatorItem
	err := json.Unmarshal([]byte(str1), &OperatorItem)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", OperatorItem)
}
