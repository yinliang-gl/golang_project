package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type AppPackageFreqFilter struct {
	needFilter      		bool
	timeLimitOfPutType1 	int64
	timeLimitOfPutType2 	int64
	currentSeconds			int64
}

func (apf *AppPackageFreqFilter) Init(tmpStr string) {
	puttypeStr := strings.Split(tmpStr, ":")
	for _, itemStr := range puttypeStr {
		subTmpStr := strings.Split(itemStr, ",")
		if len(subTmpStr) != 2 {
			continue
		}

		valInt, err := strconv.Atoi(subTmpStr[1])
		if err != nil {
			continue
		}

		if len(subTmpStr[0]) == 0 {
			continue
		}
		if subTmpStr[0][0] == '1' {
			apf.timeLimitOfPutType1 = int64(valInt)
		} else if subTmpStr[0][0] == '2' {
			apf.timeLimitOfPutType2 = int64(valInt)
		}
	}
	apf.currentSeconds = time.Now().Unix()
	// adslog.Debug("yinliang_AppPackageFreqFilter tmpStr:%#v, apf.needFilter:%#v, apf.timeLimitOfPutType1:%#v, apf.timeLimitOfPutType2:%#v", tmpStr, apf.needFilter, apf.timeLimitOfPutType1, apf.timeLimitOfPutType2)
}

func main() {
	appPackageFreqFilter := &AppPackageFreqFilter{}
	appPackageFreqFilter.Init("1,60:2,600")
	fmt.Println(appPackageFreqFilter.timeLimitOfPutType2)
}
