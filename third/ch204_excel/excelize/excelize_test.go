package excelize

import (
	"fmt"
	"testing"

	"github.com/xuri/excelize/v2"
)

var (
	records = [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}
)

func TestSimpleWrite(t *testing.T) {
	f := excelize.NewFile()

	sheetIndex := f.NewSheet("Sheet1")
	fmt.Println("add sheet index:", sheetIndex)
}
