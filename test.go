package main

import (
	"encoding/json"
	"fmt"
)

type FingerPrintJsonObj struct {
	Title  []uint64
	Video  []uint64
	Images []([]uint64)
}

func main() {
	var fingerPrintJsonObj FingerPrintJsonObj
	fingerPrintJsonObj.Title = append(fingerPrintJsonObj.Title, 321)
	fmt.Println(fingerPrintJsonObj)
}

func main1() {
	var s FingerPrintJsonObj

	str1 := "{\n" +
		"  \"title\": [\"31231312646\",\"31231312647\"],\n" +
		"  \"resources\": {\n" +
		"    \"video”: [\"41231312646\",\"43124324321\"],\n" +
		"    \"images\": [\n" +
		"      [\"51231312646\",\"51231312648\"],\n" +
		"      [\"51231312646\",\"51231312648\"],\n" +
		"      [\"51231312646\",\"51231312648\"]\n" +
		"    ]\n" +
		"  }\n" +
		"}"

	fmt.Println(len(str1))
	str := `{"title":[4231432,4321432],"video":[123,456],"images":[[123,124],[456,457]]}`
	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	for _, item := range s.Title {
		fmt.Println(item)
	}
	//fmt.Println(s.Servers[0].ServerName)

	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("JSON ERR:", err)
	}
	fmt.Println(string(b))
}
