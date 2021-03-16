package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"strconv"
	"strings"
)

type RunInfos struct {
	AdIds          []int64
	ImageStartTime string
	ImageEndTime   string
	VideoStartTime string
	VideoEndTime   string
	TitleEntices   map[int64][]TitleEntity
}

type TitleEntity struct {
	AllString map[string]string
}

func readXlsx(filename string) RunInfos {
	var supportCreativeTypeMap map[string][]int64
	supportCreativeTypeMap["广点通-横版大图&横版视频&竖版视频"] = append(supportCreativeTypeMap["广点通-横版大图&横版视频&竖版视频"], 100)

	var runInfos RunInfos
	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
	}

	fmt.Printf("length of xlFile.Sheets is %v\n", len(xlFile.Sheets))
	for _, sheet := range xlFile.Sheets {
		fmt.Printf("Sheet Name: %s\n", sheet.Name)

		var supportCreativeType []int64

		if sheet.Name == "汇总信息" {
			FillSummaryInformation(sheet, runInfos)
			continue
		} else {
			var find bool
			if supportCreativeType, find = supportCreativeTypeMap[sheet.Name]; !find {
				fmt.Errorf("不支持的类型[%v]", sheet.Name)
				continue
			}
			var titleEntitys []TitleEntity

			for i := 1; i < len(sheet.Rows); i++ {
				for j := 0; j < len(sheet.Rows[i].Cells); j++ {

				}
			}

			for _, val := range supportCreativeType {
				runInfos.TitleEntices[val] = titleEntitys
			}

		}
	}
	return runInfos
}

func FillSummaryInformation(sheet *xlsx.Sheet, runInfos RunInfos) {
	if len(sheet.Rows) != 2 {
		panic("汇总信息格式错误")
	}
	if len(sheet.Rows[1].Cells) != 5 {
		panic("汇总、" +
			"信息格式错误")
	}
	// 1. 模板广告id
	templateIdsStr := sheet.Rows[1].Cells[0].String()
	templateIdsStrSplit := strings.Split(templateIdsStr, ",")
	for _, adidStr := range templateIdsStrSplit {
		adId, err := strconv.ParseInt(adidStr, 10, 64)
		if err != nil {
			panic("模板id转为int64失败" + adidStr)
		}
		runInfos.AdIds = append(runInfos.AdIds, adId)
	}
	// 2. 图片开始结束时间
	runInfos.ImageStartTime = sheet.Rows[1].Cells[1].String()
	runInfos.ImageEndTime = sheet.Rows[1].Cells[2].String()
	// 3. 视频开始结束时间
	runInfos.ImageStartTime = sheet.Rows[1].Cells[1].String()
	runInfos.ImageEndTime = sheet.Rows[1].Cells[2].String()
}

func main() {
	excelFileName := "/Users/qtt/Documents/work_places/work_place_golang/work_tools_golang/tmp/tencentad/" +
		"postman请求-15-18814682-正保科技-中级会计师-微信/广点通_账户id_2021_01_22_15_02.xlsx"
	oraList := readXlsx(excelFileName)

	fmt.Printf("%#v", oraList)
}
