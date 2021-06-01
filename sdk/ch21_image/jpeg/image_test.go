package jpeg

import (
	"fmt"
	"image/jpeg"
	"os"
	"testing"
)

// image/jpeg: jpeg图片编解码 (功能很少)

// 获取视频宽高
func TestGetImgWH(t *testing.T) {
	file, err := os.Open("src.jpeg")
	if err != nil {
		panic(err)
	}

	imgCfg, err := jpeg.DecodeConfig(file)
	if err != nil {
		panic(err)
	}

	w := imgCfg.Width
	h := imgCfg.Height
	fmt.Printf("image width and height: %dx%d\n", w, h)
}

// jpeg图片压缩
func TestCompress(t *testing.T) {
	file, err := os.Open("src.jpeg")
	if err != nil {
		panic(err)
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		panic(err)
	}

	file, err = os.Create("tmp.jpeg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	options := &jpeg.Options{Quality: jpeg.DefaultQuality} // 图片质量(1~100,值越高越清晰)
	err = jpeg.Encode(file, img, options)
	if err != nil {
		panic(err)
	}
}
