package main

/**
号称是golang最快的json库
*/
import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)

}

func Decode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

type RtaExecuteResultCode int

const (
	RTA_SUCCESS                  RtaExecuteResultCode = 0 // 正常执行
	RTA_REQUEST_FAILED           RtaExecuteResultCode = 1 // 请求广告主服务器失败，包括超时，网络不可用等
	RTA_DEVICES_EMPTY            RtaExecuteResultCode = 2 // 设备信息为空
	RTA_ANALYSIS_RESPONSE_FAILED RtaExecuteResultCode = 3 // 解析返回体失败
	RTA_ADMIXTURE                RtaExecuteResultCode = 4 // 掺量
)

type RTAInfo struct {
	RtaID             int64
	ExcludeUserID     map[int64]bool
	ExcludeUnitID     map[int64]map[int64]bool //map[unitID]map[userID]bool
	ExTime            int64
	StatusCode        int32
	HitCache          bool                 //
	FilterFlag        bool                 //过滤开关
	ExceptionFlag     bool                 //超时或者状态码错误
	ExecuteResultCode RtaExecuteResultCode // 访问rta的结果编码
}

func main() {
	aaa := RTAInfo{RtaID: 100, ExcludeUserID: map[int64]bool{1: true}}
	byte, err := Encode(aaa)
	if err == nil {
		fmt.Println(string(byte))
	}
}
