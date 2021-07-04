package ch21_image

import (
	"image"
	"testing"
)

// image: 实现一基本的二维图像库
// 更多图像相关功能可使用第三方库: "golang.org/x/image", "github.com/chai2010/webp"

// 变量
func TestVar(t *testing.T) {
	_ = image.Black
	_ = image.White
	_ = image.Transparent
	_ = image.Opaque
}

// 错误
func TestErr(t *testing.T) {
	_ = image.ErrFormat // image: unknown format
}

// 函数
// func RegisterFormat(name, magic string, decode func(io.Reader) (Image, error), decodeConfig func(io.Reader) (Config, error)) // 注册一种图像格式的解码器

// 类型
// type Alpha：内存中的图像，At方法返回color.Alpha。
//    func NewAlpha(r Rectangle) *Alpha
//    func (p *Alpha) AlphaAt(x, y int) color.Alpha
//    func (p *Alpha) At(x, y int) color.Color
//    func (p *Alpha) Bounds() Rectangle
//    func (p *Alpha) ColorModel() color.Model
//    func (p *Alpha) Opaque() bool
//    func (p *Alpha) PixOffset(x, y int) int
//    func (p *Alpha) Set(x, y int, c color.Color)
//    func (p *Alpha) SetAlpha(x, y int, c color.Alpha)
//    func (p *Alpha) SubImage(r Rectangle) Image

// type Alpha16：内存中的图像，At方法返回color.Alpha16。
//    func NewAlpha16(r Rectangle) *Alpha16
//    func (p *Alpha16) Alpha16At(x, y int) color.Alpha16
//    func (p *Alpha16) At(x, y int) color.Color
//    func (p *Alpha16) Bounds() Rectangle
//    func (p *Alpha16) ColorModel() color.Model
//    func (p *Alpha16) Opaque() bool
//    func (p *Alpha16) PixOffset(x, y int) int
//    func (p *Alpha16) Set(x, y int, c color.Color)
//    func (p *Alpha16) SetAlpha16(x, y int, c color.Alpha16)
//    func (p *Alpha16) SubImage(r Rectangle) Image

// type CMYK：内存中的图像，At方法返回color.CMYK。
//    func NewCMYK(r Rectangle) *CMYK
//    func (p *CMYK) At(x, y int) color.Color
//    func (p *CMYK) Bounds() Rectangle
//    func (p *CMYK) CMYKAt(x, y int) color.CMYK
//    func (p *CMYK) ColorModel() color.Model
//    func (p *CMYK) Opaque() bool
//    func (p *CMYK) PixOffset(x, y int) int
//    func (p *CMYK) Set(x, y int, c color.Color)
//    func (p *CMYK) SetCMYK(x, y int, c color.CMYK)
//    func (p *CMYK) SubImage(r Rectangle) Image

// type Config：图像颜色模式和尺寸
//    func DecodeConfig(r io.Reader) (Config, string, error)

// type Gray：内存中的图像，At方法返回color.GRAY。
//    func NewGray(r Rectangle) *Gray
//    func (p *Gray) At(x, y int) color.Color
//    func (p *Gray) Bounds() Rectangle
//    func (p *Gray) ColorModel() color.Model
//    func (p *Gray) GrayAt(x, y int) color.Gray
//    func (p *Gray) Opaque() bool
//    func (p *Gray) PixOffset(x, y int) int
//    func (p *Gray) Set(x, y int, c color.Color)
//    func (p *Gray) SetGray(x, y int, c color.Gray)
//    func (p *Gray) SubImage(r Rectangle) Image

// type Gray16：内存中的图像，At方法返回color.GRAY16。
//    func NewGray16(r Rectangle) *Gray16
//    func (p *Gray16) At(x, y int) color.Color
//    func (p *Gray16) Bounds() Rectangle
//    func (p *Gray16) ColorModel() color.Model
//    func (p *Gray16) Gray16At(x, y int) color.Gray16
//    func (p *Gray16) Opaque() bool
//    func (p *Gray16) PixOffset(x, y int) int
//    func (p *Gray16) Set(x, y int, c color.Color)
//    func (p *Gray16) SetGray16(x, y int, c color.Gray16)
//    func (p *Gray16) SubImage(r Rectangle) Image

// type Image：图像是一个有限的矩形网格的color.Color。颜色值取自颜色模型。
//    func Decode(r io.Reader) (Image, string, error)

// type NRGBA：内存中的图像，At方法返回color.NRGBA。
//    func NewNRGBA(r Rectangle) *NRGBA
//    func (p *NRGBA) At(x, y int) color.Color
//    func (p *NRGBA) Bounds() Rectangle
//    func (p *NRGBA) ColorModel() color.Model
//    func (p *NRGBA) NRGBAAt(x, y int) color.NRGBA
//    func (p *NRGBA) Opaque() bool
//    func (p *NRGBA) PixOffset(x, y int) int
//    func (p *NRGBA) Set(x, y int, c color.Color)
//    func (p *NRGBA) SetNRGBA(x, y int, c color.NRGBA)
//    func (p *NRGBA) SubImage(r Rectangle) Image

// type NRGBA64：内存中的图像，At方法返回color.NRGBA64。
//    func NewNRGBA64(r Rectangle) *NRGBA64
//    func (p *NRGBA64) At(x, y int) color.Color
//    func (p *NRGBA64) Bounds() Rectangle
//    func (p *NRGBA64) ColorModel() color.Model
//    func (p *NRGBA64) NRGBA64At(x, y int) color.NRGBA64
//    func (p *NRGBA64) Opaque() bool
//    func (p *NRGBA64) PixOffset(x, y int) int
//    func (p *NRGBA64) Set(x, y int, c color.Color)
//    func (p *NRGBA64) SetNRGBA64(x, y int, c color.NRGBA64)
//    func (p *NRGBA64) SubImage(r Rectangle) Image

// type NYCbCrA：是非 alpha 预乘 Y'CbCr 与 alpha 颜色的内存图像。 A 和 ASride 类似于嵌入式 YCbCr 的 Y 和 YStride 字段。
//    func NewNYCbCrA(r Rectangle, subsampleRatio YCbCrSubsampleRatio) *NYCbCrA
//    func (p *NYCbCrA) AOffset(x, y int) int
//    func (p *NYCbCrA) At(x, y int) color.Color
//    func (p *NYCbCrA) ColorModel() color.Model
//    func (p *NYCbCrA) NYCbCrAAt(x, y int) color.NYCbCrA
//    func (p *NYCbCrA) Opaque() bool
//    func (p *NYCbCrA) SubImage(r Rectangle) Image

// type Paletted：给定调色板的uint8索引的内存图像
//    func NewPaletted(r Rectangle, p color.Palette) *Paletted
//    func (p *Paletted) At(x, y int) color.Color
//    func (p *Paletted) Bounds() Rectangle
//    func (p *Paletted) ColorIndexAt(x, y int) uint8
//    func (p *Paletted) ColorModel() color.Model
//    func (p *Paletted) Opaque() bool
//    func (p *Paletted) PixOffset(x, y int) int
//    func (p *Paletted) Set(x, y int, c color.Color)
//    func (p *Paletted) SetColorIndex(x, y int, index uint8)
//    func (p *Paletted) SubImage(r Rectangle) Image

// type PalettedImage：一种图像，颜色可能来自于有限调色板。

// type Point：点是 X、Y 坐标对。轴向右和向下增加。
//    func Pt(X, Y int) Point
//    func (p Point) Add(q Point) Point
//    func (p Point) Div(k int) Point
//    func (p Point) Eq(q Point) bool
//    func (p Point) In(r Rectangle) bool
//    func (p Point) Mod(r Rectangle) Point
//    func (p Point) Mul(k int) Point
//    func (p Point) String() string
//    func (p Point) Sub(q Point) Point

// type RGBA：内存中的图像，At方法返回color.RGBA。
//    func NewRGBA(r Rectangle) *RGBA
//    func (p *RGBA) At(x, y int) color.Color
//    func (p *RGBA) Bounds() Rectangle
//    func (p *RGBA) ColorModel() color.Model
//    func (p *RGBA) Opaque() bool
//    func (p *RGBA) PixOffset(x, y int) int
//    func (p *RGBA) RGBAAt(x, y int) color.RGBA
//    func (p *RGBA) Set(x, y int, c color.Color)
//    func (p *RGBA) SetRGBA(x, y int, c color.RGBA)
//    func (p *RGBA) SubImage(r Rectangle) Image

// type RGBA64：内存中的图像，At方法返回color.RGBA64。
//    func NewRGBA64(r Rectangle) *RGBA64
//    func (p *RGBA64) At(x, y int) color.Color
//    func (p *RGBA64) Bounds() Rectangle
//    func (p *RGBA64) ColorModel() color.Model
//    func (p *RGBA64) Opaque() bool
//    func (p *RGBA64) PixOffset(x, y int) int
//    func (p *RGBA64) RGBA64At(x, y int) color.RGBA64
//    func (p *RGBA64) Set(x, y int, c color.Color)
//    func (p *RGBA64) SetRGBA64(x, y int, c color.RGBA64)
//    func (p *RGBA64) SubImage(r Rectangle) Image

// type Rectangle：矩形
//    func Rect(x0, y0, x1, y1 int) Rectangle
//    func (r Rectangle) Add(p Point) Rectangle
//    func (r Rectangle) At(x, y int) color.Color
//    func (r Rectangle) Bounds() Rectangle
//    func (r Rectangle) Canon() Rectangle
//    func (r Rectangle) ColorModel() color.Model
//    func (r Rectangle) Dx() int
//    func (r Rectangle) Dy() int
//    func (r Rectangle) Empty() bool
//    func (r Rectangle) Eq(s Rectangle) bool
//    func (r Rectangle) In(s Rectangle) bool
//    func (r Rectangle) Inset(n int) Rectangle
//    func (r Rectangle) Intersect(s Rectangle) Rectangle
//    func (r Rectangle) Overlaps(s Rectangle) bool
//    func (r Rectangle) Size() Point
//    func (r Rectangle) String() string
//    func (r Rectangle) Sub(p Point) Rectangle
//    func (r Rectangle) Union(s Rectangle) Rectangle

// type Uniform：统一是一个无限大小的统一颜色的图像。它实现了 color.Color、color.Model 和 Image 接口。
//    func NewUniform(c color.Color) *Uniform
//    func (c *Uniform) At(x, y int) color.Color
//    func (c *Uniform) Bounds() Rectangle
//    func (c *Uniform) ColorModel() color.Model
//    func (c *Uniform) Convert(color.Color) color.Color
//    func (c *Uniform) Opaque() bool
//    func (c *Uniform) RGBA() (r, g, b, a uint32)

// type YCbCr：Y'CbCr颜色的内存图像。
//    func NewYCbCr(r Rectangle, subsampleRatio YCbCrSubsampleRatio) *YCbCr
//    func (p *YCbCr) At(x, y int) color.Color
//    func (p *YCbCr) Bounds() Rectangle
//    func (p *YCbCr) COffset(x, y int) int
//    func (p *YCbCr) ColorModel() color.Model
//    func (p *YCbCr) Opaque() bool
//    func (p *YCbCr) SubImage(r Rectangle) Image
//    func (p *YCbCr) YCbCrAt(x, y int) color.YCbCr
//    func (p *YCbCr) YOffset(x, y int) int

// type YCbCrSubsampleRatio：YCbCr 图像中使用的色度子样本比率。
//    func (s YCbCrSubsampleRatio) String() string
