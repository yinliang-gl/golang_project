package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

type RunInfos struct {
	ImageStartTime string
	ImageEndTime   string
	VideoStartTime string
	VideoEndTime   string
	TitleEntices   map[int64][]TitleEntity
}

type TitleEntity struct {
	AllString map[string][]string
}

func readXlsx(filename string) []RunInfos {
	var listOra []RunInfos
	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
	}
	fmt.Printf("length of xlFile.Sheets is %v\n", len(xlFile.Sheets))
	for _, sheet := range xlFile.Sheets {
		fmt.Printf("Sheet Name: %s\n", sheet.Name)
		tmpOra := RunInfos{}
		for _, row := range sheet.Rows {
			var strs []string
			for _, cell := range row.Cells {
				text := cell.String()
				strs = append(strs, text)
			}
			fmt.Println("================================")
			fmt.Println(sheet.Name)
			fmt.Println(strs)
			listOra = append(listOra, tmpOra)
		}
	}
	return listOra
}

func main() {
	excelFileName := "/Users/qtt/Documents/work_places/work_place_golang/work_tools_golang/tmp/tencentad/" +
		"postman请求-15-18814682-正保科技-中级会计师-微信/广点通_账户id_2021_01_22_15_02.xlsx"
	oraList := readXlsx(excelFileName)

	fmt.Println("%#v", oraList)
}
