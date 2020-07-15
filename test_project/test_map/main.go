package main

import (
	"fmt"
	"strconv"
	"time"
	"testing"
)

type Key struct {
	MediaId      string
	AdId, SlotId int
}

type KeyV2 struct {
	MediaId      int
	AdId, SlotId int
}

func TestTime01(t *testing.T) {

}

func TestStructMap(t *testing.T) {
	structV1Map := make(map[Key]bool)
	structV2Map := make(map[KeyV2]bool)
	strintMap := make(map[string]bool)

	// stringArr := make([]string)

	var stringArr []string

	COUNT := 1000000
	for i := 0; i < COUNT; i++ {
		stringArr = append(stringArr, strconv.Itoa(i)+"21312312312")
		structV1Key := Key{MediaId: strconv.Itoa(i), SlotId: i, AdId: i}
		structV2Key := KeyV2{MediaId: i, SlotId: i, AdId: i}
		stringkey := strconv.Itoa(i) + "_" + strconv.Itoa(i) + "_" + strconv.Itoa(i)
		if (i % 2) == 0 {
			structV1Map[structV1Key] = true
			structV2Map[structV2Key] = true
			strintMap[stringkey] = true
		} else {
			structV1Map[structV1Key] = false
			structV2Map[structV2Key] = false
			strintMap[stringkey] = false
		}
	}

	start := time.Now().UnixNano() / 1e6
	for i := 0; i < COUNT; i++ {
		structKey := Key{MediaId: strconv.Itoa(i), SlotId: i, AdId: i}
		if _, find := structV1Map[structKey]; find {
			// fmt.Printf("%#v-%#v\n", i, val)
		}
	}
	fmt.Printf("struct_key_v1(有数字转字符串)  %dms\n", time.Now().UnixNano()/1e6-start)

	// return

	start = time.Now().UnixNano() / 1e6
	for i := 0; i < COUNT; i++ {
		structKey := KeyV2{MediaId: i, SlotId: i, AdId: i}
		if _, find := structV2Map[structKey]; find {
			// fmt.Printf("%#v-%#v\n", i, val)
		}
	}
	fmt.Printf("struct_key_v2(没有数字转字符串)  %dms\n", time.Now().UnixNano()/1e6-start)

	start = time.Now().UnixNano() / 1e6
	for i := 0; i < COUNT; i++ {
		stringkey := strconv.Itoa(i) + "_" + strconv.Itoa(i) + "_" + strconv.Itoa(i)
		if _, find := strintMap[stringkey]; find {
			//fmt.Printf("%#v-%#v\n", i, val)
		}
	}
	fmt.Printf("string_key_v1(有数字转字符串)  %dms\n", time.Now().UnixNano()/1e6-start)

	start = time.Now().UnixNano() / 1e6
	for i := 0; i < COUNT; i++ {
		stringkey := stringArr[i] + "_" + stringArr[i] + "_" + stringArr[i]
		if _, find := strintMap[stringkey]; find {
			//fmt.Printf("%#v-%#v\n", i, val)
		}
	}
	fmt.Printf("string_key_v2(没有数字转字符串)  %dms\n", time.Now().UnixNano()/1e6-start)
}
