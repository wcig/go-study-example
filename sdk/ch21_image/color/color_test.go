package color

import (
	"image/color"
	"testing"
)

// image/color：实现一基础的颜色库

// 变量
func TestVar(t *testing.T) {
	_ = color.Black
	_ = color.White
	_ = color.Transparent
	_ = color.Opaque
}

// 函数
// func CMYKToRGB(c, m, y, k uint8) (uint8, uint8, uint8)     // CMYK四元组->RGB三元组
// func RGBToCMYK(r, g, b uint8) (uint8, uint8, uint8, uint8) // RGB三元组->CMYK四元组
// func RGBToYCbCr(r, g, b uint8) (uint8, uint8, uint8)       // RGB三元组->Y'CbCr三元组
// func YCbCrToRGB(y, cb, cr uint8) (uint8, uint8, uint8)     // Y'CbCr三元组->RGB三元组

// 类型
// 1.Alpha：8bit Alpha颜色
// type Alpha struct {
//    A uint8
// }
// func (c Alpha) RGBA() (r, g, b, a uint32)

// 2.Alpha16：16bit Alpha颜色
// type Alpha16 struct {
// 	A uint16
// }
// func (c Alpha16) RGBA() (r, g, b, a uint32)

// 3.CMYK：完全不透明的CMYK颜色，青、品红、黄、黑各有8bit
// type CMYK struct {
// 	C, M, Y, K uint8
// }
// func (c CMYK) RGBA() (uint32, uint32, uint32, uint32)

// 4.Color：可以转换自身到alpha预乘每通道16bitRGBA，转换可能有损失
// type Color interface {
// 	RGBA() (r, g, b, a uint32)
// }

// 5.Gray：8bit灰度颜色
// type Gray struct {
// 	Y uint8
// }
//
// func (c Gray) RGBA() (r, g, b, a uint32)

// 6.Gray16：16bit灰度颜色
// type Gray16 struct {
// 	Y uint16
// }
// func (c Gray16) RGBA() (r, g, b, a uint32)

// 7.Model：将任何颜色从自身颜色模式转换诶其他人也颜色
// type Model interface {
// 	Convert(c Color) Color
// }
// var (
// 	RGBAModel    Model = ModelFunc(rgbaModel)
// 	RGBA64Model  Model = ModelFunc(rgba64Model)
// 	NRGBAModel   Model = ModelFunc(nrgbaModel)
// 	NRGBA64Model Model = ModelFunc(nrgba64Model)
// 	AlphaModel   Model = ModelFunc(alphaModel)
// 	Alpha16Model Model = ModelFunc(alpha16Model)
// 	GrayModel    Model = ModelFunc(grayModel)
// 	Gray16Model  Model = ModelFunc(gray16Model)
// )
// func ModelFunc(f func(Color) Color) Model // 返回调用f实现转换的模型

// 8.NRGBA：非 alpha 预乘的 32 位颜色
// type NRGBA struct {
// 	R, G, B, A uint8
// }
// func (c NRGBA) RGBA() (r, g, b, a uint32)

// 9.NRGBA64：非 alpha 预乘的 64 位颜色，红色、绿色、蓝色和 alpha 各有 16 位
// type NRGBA64 struct {
// 	R, G, B, A uint16
// }
// func (c NRGBA64) RGBA() (r, g, b, a uint32)

// 10.NYCbCrA：一种非 alpha 预乘 Y'CbCr-with-alpha 颜色，每个颜色有 8 位，用于一个亮度、两个色度和一个 alpha 分量
// type NYCbCrA struct {
// 	YCbCr
// 	A uint8
// }
// func (c NYCbCrA) RGBA() (uint32, uint32, uint32, uint32)

// 11.Palette：颜色调色板
// type Palette []Color
// func (p Palette) Convert(c Color) Color // 返回在欧几里得 R、G、B 空间中最接近 c 的调色板颜色
// func (p Palette) Index(c Color) int     // 返回欧几里得 R、G、B、A 空间中最接近 c 的调色板颜色的索引

// 12.RGBA：传统的 32 位 alpha 预乘颜色，红色、绿色、蓝色和 alpha 各有 8 位
// type RGBA struct {
// 	R, G, B, A uint8
// }
// func (c RGBA) RGBA() (r, g, b, a uint32)

// 13.RGBA64：64 位 alpha 预乘颜色，红色、绿色、蓝色和 alpha 各有 16 位
// type RGBA64 struct {
// 	R, G, B, A uint16
// }
// func (c RGBA64) RGBA() (r, g, b, a uint32)

// 14.YCbCr：一种完全不透明的 24 位 Y'CbCr 颜色，每个 8 位用于一个亮度和两个色度分量
// type YCbCr struct {
// 	Y, Cb, Cr uint8
// }
// func (c YCbCr) RGBA() (uint32, uint32, uint32, uint32)
