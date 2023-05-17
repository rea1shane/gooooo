package excel

import (
	"github.com/xuri/excelize/v2"
	"strconv"
	"testing"
)

func TestNewExcel(t *testing.T) {
	f := excelize.NewFile()
	defer func() {
		f.Close()
	}()

	data := make(map[string][]int)
	data["小红"] = []int{1, 2, 3}
	data["小明"] = []int{4, 5, 6, 7}
	data["小刚"] = []int{8}

	sheetName := "A"
	f.NewSheet(sheetName)
	y := 1
	for s, ints := range data {
		startY := y
		f.SetCellValue(sheetName, "A"+strconv.Itoa(y), s)
		for _, i := range ints {
			f.SetCellValue(sheetName, "B"+strconv.Itoa(y), i)
			y++
		}
		f.MergeCell(sheetName, "A"+strconv.Itoa(startY), "A"+strconv.Itoa(y-1))
	}

	f.DeleteSheet("Sheet1")
	err := f.SaveAs("generate_by_go.xlsx")
	if err != nil {
		panic(err)
	}
}
