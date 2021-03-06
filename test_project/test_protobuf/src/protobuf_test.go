package src

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/yinliang-gl/golang_project/test_project/test_protobuf/proto"
	"log"
	"testing"
)

func TestProtobuf001(t *testing.T) {
	score_info := &ptoto.BaseScoreInfoT{}
	score_info.WinCount = new(int32)
	*score_info.WinCount = 1
	score_info.LoseCount = new(int32)
	*score_info.LoseCount = 2
	score_info.ExceptionCount = new(int32)
	*score_info.ExceptionCount = 3
	score_info.KillCount = new(int32)
	*score_info.KillCount = 4
	score_info.DeathCount = new(int32)
	*score_info.DeathCount = 5
	score_info.AssistCount = new(int32)
	*score_info.AssistCount = 6
	score_info.Rating = new(int64)
	*score_info.Rating = 1800

	fmt.Printf("original data{%s}\n", score_info.String())

	// encode
	data, err := proto.Marshal(score_info)
	if err != nil {
		fmt.Printf("proto encode error[%s]\n", err.Error())
		return
	}

	// decode
	score_info_1 := &ptoto.BaseScoreInfoT{}
	err = proto.Unmarshal(data, score_info_1)
	if err != nil {
		fmt.Printf("proto decode error[%s]\n", err.Error())
		return
	}

	*score_info_1.Rating = 2000
	fmt.Printf("after decode:{%s}\n", score_info_1.String())

	result, err := json.MarshalIndent(score_info_1, "", "    ")
	if err != nil {
		log.Fatal("GameData arshalIndent err: " + err.Error() + "\n")
	}

	fmt.Println(string(result))

}
