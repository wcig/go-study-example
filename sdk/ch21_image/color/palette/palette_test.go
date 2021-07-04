package palette

import (
	"image/color/palette"
	"testing"
)

// image/color/palette：提供标准颜色调色板

// 变量
func TestVar(t *testing.T) {
	_ = palette.Plan9
	_ = palette.WebSafe
}
