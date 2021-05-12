package ch31_xls

import (
	"os"
	"testing"

	"github.com/tealeg/xlsx/v3"
)

// github.com/tealeg/xlsx工具包处理xls

// 简单使用: 内存占用偏高,大数据量情况下慎用 (本示例最高占用3.839g)
func TestTealegSimpleWrite(t *testing.T) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		panic(err)
	}

	for i := 1; i <= 200000; i++ {
		row := sheet.AddRow()
		for j := 1; j <= 30; j++ {
			cell := row.AddCell()
			cell.Value = "1234567890123456789012345678901234567890"
		}
	}

	name := "test-tealeg.xlsx"
	if err := file.Save(name); err != nil {
		panic(err)
	}
	_ = os.Remove(name)
}
