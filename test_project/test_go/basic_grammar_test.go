package test_go

import (
	"encoding/json"
	"fmt"
	"testing"
)

/**
测试三个点的用法
*/
func TestBasicGrammar001(t *testing.T) {
	var strss = []string{
		"qwr",
		"234",
		"yui",
	}
	var strss2 = []string{
		"qqq",
		"aaa",
		"zzz",
		"zzz",
	}
	strss = append(strss, strss2...) //strss2的元素被打散一个个append进strss
	fmt.Println(strss)
}
func TestBasicGrammar002(t *testing.T) {
	var strss = []string{
		"qwr",
		"234",
		"yui",
		"cvbc",
	}

	iterater_fun := func(args ...string) { //可以接受任意个string参数
		for _, v := range args {
			fmt.Println(v)
		}
	}

	iterater_fun(strss...) //切片被打散传入
}

/**
测试将类json化
*/
type Order struct {
	Cid        string `json:"cid" form:"cid" superChecker:"cid"`
	Pversion   string `json:"pversion" form:"pversion" superChecker:"pversion"`
	Phone      string `json:"phone" form:"phone" superChecker:"mobilephone|telephone"`
	Uname      string `json:"uname" form:"uname" superChecker:"uname"`
	Money      string `json:"money" form:"money" superChecker:"money"`
	BillType   string `json:"bill_type" form:"bill_type" superChecker:"bill_type"`
	ProductId  string `json:"product_id" form:"product_id" superChecker:"product_id"`
	TravelDate string `json:"travel_date" form:"travel_date" superChecker:"travel_date"`
	AppType    string `json:"app_type" form:"app_type" superChecker:"app_type"`
	Num        string `json:"num" form:"num" superChecker:"num"`
	OpenBy     string `json:"open_by" form:"open_by"`
	UserId     string `json:"zonst_user_id" form:"zonst_user_id"`
	Addr       string `json:"addr" from:"addr"`
}

/**
测试将类json化
*/
func TestBasicGrammar003(t *testing.T) {
	orderNew := Order{
		Cid:       "20030",
		Pversion:  "v1.0",
		Phone:     "18978376478",
		Uname:     "石勇,张三,赵四",
		Money:     "10080",
		BillType:  "1",
		ProductId: "1005",
		//TravelDate :"2018-05-24",
		Num:     "3",
		AppType: "1",
		OpenBy:  "2",
		UserId:  "1111",
		Addr:    "江西南昌",
	}
	buf, err := json.Marshal(orderNew)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
}
