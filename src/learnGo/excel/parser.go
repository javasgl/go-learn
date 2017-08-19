package main

import (
	"fmt"
	"strconv"

	"github.com/Luxurioust/excelize"
)

const (
	XLSX_FILE_SUPER_STAR = "/Users/songgl/Desktop/super-star.xlsx"
	XLSX_FILE_INS        = "/Users/songgl/Desktop/ins-tv.xlsx"
)

func main() {

	parse(XLSX_FILE_INS)
	// parse(XLSX_FILE_SUPER_STAR)

}

func parse(file string) {

	xlsx, err := excelize.OpenFile(file)

	if err != nil {
		fmt.Println(err)
	}

	sheets := xlsx.GetSheetMap()

	fmt.Println(sheets)

	for index, _ := range sheets {
		rows := xlsx.GetRows("sheet" + strconv.Itoa(index))
		for _, row := range rows {
			for _, colCell := range row {
				fmt.Print(colCell, "\t")
			}
			fmt.Println()
		}
	}

}
