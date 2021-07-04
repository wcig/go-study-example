package png

// image/png：png图像编解码（png规范见：https://www.w3.org/TR/PNG/）

// 函数
// func Decode(r io.Reader) (image.Image, error) // 从r读取一png图像，解析并返回image.Image和error
// func DecodeConfig(r io.Reader) (image.Config, error) // 从r读取一png图像，解析并返回image.Config和error
// func Encode(w io.Writer, m image.Image, o *Options) error // 使用给定的选项以 png 4:2:0 基线格式将图像 m 写入 w。如果传递了 nil *Options，则使用默认参数。

// 类型
// 1.CompressionLevel：压缩级别
// type CompressionLevel int
// const (
//    DefaultCompression CompressionLevel = 0
//    NoCompression      CompressionLevel = -1
//    BestSpeed          CompressionLevel = -2
//    BestCompression    CompressionLevel = -3
// )

// 2.Encoder：png编码器
// type Encoder struct {
//    CompressionLevel CompressionLevel
//
//    // BufferPool optionally specifies a buffer pool to get temporary
//    // EncoderBuffers when encoding an image.
//    BufferPool EncoderBufferPool // Go 1.9
// }
// func (enc *Encoder) Encode(w io.Writer, m image.Image) error // 以PNG格式写入图像到w

// 3.EncoderBuffer：带缓冲区的PNG编码器
// type EncoderBuffer encoder

// 4.EncoderBufferPool：是一个用于获取和返回 EncoderBuffer 结构的临时实例的接口。这可用于在编码多个图像时重用缓冲区。
// type EncoderBufferPool interface {
//    Get() *EncoderBuffer
//    Put(*EncoderBuffer)
// }

// 5.FormatError：非JPEG输入的格式错误
// type FormatError string
// func (e FormatError) Error() string

// 6.UnsupportedError：// type Reader
// type UnsupportedError string
// func (e UnsupportedError) Error() string
