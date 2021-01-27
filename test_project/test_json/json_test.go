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

type SetInfo struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	ErrMsg string `json:"err_msg"`
}

type ConsistencyInfo struct {
	AccountId  int64     `json:"account_id"`
	CampaignId int64     `json:"campaign_id"`
	AdGroupId  int64     `json:"ad_group_id"`
	SetInfos   []SetInfo `json:"set_infos"`
}

type ConsistencyInfoList struct {
	List []ConsistencyInfo `json:"list"`
}

func TestAnalyseJsonMap(t *testing.T) {
	operatorInfo := ConsistencyInfo{
		AccountId:  1,
		CampaignId: 2,
		AdGroupId:  3,
	}

	operatorInfo.SetInfos = append(operatorInfo.SetInfos, SetInfo{
		Key:    "budget",
		Value:  "100",
		ErrMsg: "can't ping",
	})

	operatorInfo.SetInfos = append(operatorInfo.SetInfos, SetInfo{
		Key:    "name",
		Value:  "new_fdsafdsfsd",
		ErrMsg: "can't ping",
	})

	operatorInfoList := ConsistencyInfoList{}
	operatorInfoList.List = append(operatorInfoList.List, operatorInfo)

	resByte, err := json.Marshal(operatorInfoList)
	if err != nil {
		fmt.Printf("error exit %#v\n", err)
	}
	fmt.Println(string(resByte))
}
