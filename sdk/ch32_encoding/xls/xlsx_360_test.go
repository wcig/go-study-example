package xls

import (
	"fmt"
	"os"
	"testing"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// github.com/360EntSecGroup-Skylar/excelize工具包处理xls (v2.4.0需要go1.15)

// 简单使用: 内存占用偏高,大数据量情况下慎用 (本示例最高占用2.791g)
func Test360SimpleWrite(t *testing.T) {
	file := excelize.NewFile()
	index := file.NewSheet("Sheet1")
	fmt.Println("add sheet index:", index)

	for i := 1; i <= 200000; i++ {
		for j := 1; j <= 30; j++ {
			axis, err := excelize.CoordinatesToCellName(j, i)
			if err != nil {
				panic(err)
			}
			value := "1234567890123456789012345678901234567890"
			if err := file.SetCellValue("Sheet1", axis, value); err != nil {
				panic(err)
			}
		}
	}

	name := "test-360.xlsx"
	if err := file.SaveAs(name); err != nil {
		panic(err)
	}
	_ = os.Remove(name)
}

// stream流写入,内存占用很少 (本示例最高占用60m)
func Test360StreamWrite(t *testing.T) {
	file := excelize.NewFile()
	index := file.NewSheet("Sheet1")
	fmt.Println("add sheet index:", index)
	streamWriter, err := file.NewStreamWriter("Sheet1")
	if err != nil {
		panic(err)
	}

	for i := 1; i <= 200000; i++ {
		values := make([]interface{}, 30)
		for j := 1; j <= 30; j++ {
			values[j-1] = "1234567890123456789012345678901234567890"
		}
		axis, err := excelize.CoordinatesToCellName(1, i)
		if err != nil {
			panic(err)
		}
		if err := streamWriter.SetRow(axis, values); err != nil {
			panic(err)
		}
	}

	if err := streamWriter.Flush(); err != nil {
		panic(err)
	}
	name := "test-360.xlsx"
	if err := file.SaveAs(name); err != nil {
		panic(err)
	}
	_ = os.Remove(name)
}
