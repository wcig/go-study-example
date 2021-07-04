package gif

import (
	"image/gif"
	"testing"
)

// image/gif：gif图像的编解码

// 常量
func TestConst(t *testing.T) {
	_ = gif.DisposalNone
	_ = gif.DisposalBackground
	_ = gif.DisposalPrevious
}

// 函数
// func Decode(r io.Reader) (image.Image, error)             // 从r读取GIF图像，返回第一个嵌入的图像作为image.Image和error
// func DecodeConfig(r io.Reader) (image.Config, error)      // 从r读取GIF图像，在不对整个图像解码的情况下获取图像的全局颜色模型和尺寸
// func Encode(w io.Writer, m image.Image, o *Options) error // 以GIF格式写入图像到w
// func EncodeAll(w io.Writer, g *GIF) error                 // 将g中的图像以GIF格式写入w，并具有给定的循环技术和帧之间延迟
// func DecodeAll(r io.Reader) (*GIF, error)                 // 从r读取GIF图像，返回顺序帧和时序信息

// 类型
// 1.GIF：可能存储多个图像在一个GIF文件中
// type GIF struct {
//    Image []*image.Paletted // The successive images.
//    Delay []int             // The successive delay times, one per frame, in 100ths of a second.
//    LoopCount int
//    Disposal []byte // Go 1.5
//    Config image.Config // Go 1.5
//    BackgroundIndex byte // Go 1.5
// }

//  2.Options：编码参数
// type Options struct {
//    NumColors int
//    Quantizer draw.Quantizer
//    Drawer draw.Drawer
// }
