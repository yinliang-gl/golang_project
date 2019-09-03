package main

import (
	"fmt"
	"git.qutoutiao.net/CPC/data_bus/tools/golang_tools/proto"
	"git.qutoutiao.net/CPC/data_bus/tools/golang_tools/util"
	"github.com/gogo/protobuf/proto"
	"runtime"
)

func main() {
	benchmark_file_util := data_bus.BenchmarkFileUtil{}
	if !benchmark_file_util.Init("/Users/qtt/Desktop/tmp/material.data") {
		fmt.Errorf("Init error")
	}

	for {
		if success, result := benchmark_file_util.Next(); success == true {
			var material model_algorithm.Material
			proto.Unmarshal([]byte(result), &material)
			//fmt.Println(material.String())
			//fmt.Println(result)
		} else {
			break
		}
	}

	fmt.Printf("\n\n\nGetTotalItemCount[%d]\n", benchmark_file_util.GetTotalItemCount())

	runtime.GC()
}
