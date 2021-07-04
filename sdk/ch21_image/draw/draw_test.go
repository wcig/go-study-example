package draw

// image/draw：提供图像合成功能。（使用可参考：https://golang.org/doc/articles/image_draw.html）

// 函数
// func Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point, op Op) // 调用nil mask的DrawMask
// func DrawMask(dst Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, op Op) // 将 dst 中的 r.Min 与 src 中的 sp 和 mask 中的 mp 对齐，然后用 Porter-Duff 组合的结果替换 dst 中的矩形 r。零掩码被视为不透明。

// 类型
// 1.Drawer：包含一Draw方法的接口
// type Drawer interface {
//    Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point)
// }

// 2.Image：包含image.Image和改变单个像素Set方法的接口
// type Image interface {
//    image.Image
//    Set(x, y int, c color.Color)
// }

// 3.Op：Porter-Duff 合成算子
// type Op
// func (op Op) Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point) // 使用此 Op 调用 Draw 函数来实现 Drawer 接口

// 4.Quantizer：图像生成调色板
// type Quantizer interface {
//    Quantize(p color.Palette, m image.Image) color.Palette
// }
