package main

import (
	"fmt"
	"strconv"
	"time"
)

var TOTAL_NUM int = 100000000

func func1() {

	elementMap := make(map[string]int)
	slotIdStr := "999999999"
	groupIdStr := "88888888"
	for i := 0; i < TOTAL_NUM; i++ {
		elementMap[slotIdStr+groupIdStr+strconv.Itoa(i)] = 1
	}

	var arr []string
	for i := 0; i < TOTAL_NUM; i++ {
		arr = append(arr, strconv.Itoa(i))
	}
	bT := time.Now() // 开始时间
	for i := 0; i < TOTAL_NUM; i++ {
		_, find := elementMap[slotIdStr+groupIdStr+arr[i%TOTAL_NUM]]
		if find {

		}
	}
	eT := time.Since(bT) // 从开始到当前所消耗的时间
	fmt.Println("Run time: ", eT)
}

type DictV2Key struct {
	slotIdStr  int
	groupIdStr int
	valInt     int
}

func func2() {

	elementMap := make(map[DictV2Key]int)
	for i := 0; i < TOTAL_NUM; i++ {
		tmpDictKey := DictV2Key{
			slotIdStr:  999999999,
			groupIdStr: 88888888,
			valInt:     i,
		}
		elementMap[tmpDictKey] = 1
	}

	var arr []int
	for i := 0; i < TOTAL_NUM; i++ {
		arr = append(arr, i)
	}
	bT := time.Now() // 开始时间
	for i := 0; i < TOTAL_NUM; i++ {
		tmpDictKey := DictV2Key{
			slotIdStr:  i * 999,
			groupIdStr: i * 88,
			valInt:     arr[i%TOTAL_NUM],
		}
		_, find := elementMap[tmpDictKey]
		if find {

		}
	}

	eT := time.Since(bT) // 从开始到当前所消耗的时间
	fmt.Println("Run time: ", eT)
}

type DictV3Key struct {
	ElementString [3]string
	ElementInt    [3]int64
}

func (this *DictV3Key) Clear() {
	var tmpStr [3]string
	this.ElementString = tmpStr

	var tmpInt [3]int64
	this.ElementInt = tmpInt
}
func func3() {
	elementMap := make(map[DictV3Key]int)
	for i := 0; i < TOTAL_NUM; i++ {
		tmpDictKey := DictV3Key{
			// ElementString: [3]string{slotIdStr, groupIdStr},
			ElementInt: [3]int64{int64(i), 3, 3},
		}
		elementMap[tmpDictKey] = 1
	}

	bT := time.Now() // 开始时间
	for i := 0; i < TOTAL_NUM; i++ {
		tmpDictKey := DictV3Key{
			// ElementString: [3]string{slotIdStr, groupIdStr},
			ElementInt: [3]int64{int64(i), 3, 3},
		}
		_, find := elementMap[tmpDictKey]
		if find {

		}
	}

	eT := time.Since(bT) // 从开始到当前所消耗的时间
	fmt.Println("Run time: ", eT)
}

func main() {
	func1()
	func2()
	func3()
}
