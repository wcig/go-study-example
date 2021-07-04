package jpeg

import (
	"fmt"
	"image/jpeg"
	"testing"
)

// image/jpeg: jpeg图片编解码 (功能简单)

// 常量
func TestConst(t *testing.T) {
	fmt.Println(jpeg.DefaultQuality) // 默认质量编码参数（75）
}

// 函数
// func Decode(r io.Reader) (image.Image, error) // 从r读取一jpeg图像，解析并返回image.Image和error
// func DecodeConfig(r io.Reader) (image.Config, error) // 从r读取一jpeg图像，解析并返回image.Config和error
// func Encode(w io.Writer, m image.Image, o *Options) error // 使用给定的选项以 JPEG 4:2:0 基线格式将图像 m 写入 w。如果传递了 nil *Options，则使用默认参数。

// 类型
// 1.FormatError：非JPEG输入的格式错误
// type FormatError string
// func (e FormatError) Error() string

// 2.Options：编码质量参数，取值范围从1到100，越高质量越好
// type Options struct {
//    Quality int
// }

// 3.Reader：弃用
// type Reader

// 4.UnsupportedError：// type Reader
// type UnsupportedError string
// func (e UnsupportedError) Error() string
