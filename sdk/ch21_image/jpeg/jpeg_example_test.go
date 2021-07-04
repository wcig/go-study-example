package jpeg

import (
	"fmt"
	"image/jpeg"
	"os"
	"testing"
)

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
	srcFile, err := os.Open("src.jpeg")
	if err != nil {
		panic(err)
	}

	img, err := jpeg.Decode(srcFile)
	if err != nil {
		panic(err)
	}

	dstFile, err := os.Create("tmp.jpeg")
	if err != nil {
		panic(err)
	}
	defer dstFile.Close()

	options := &jpeg.Options{Quality: jpeg.DefaultQuality} // 图片质量(1~100,值越高越清晰)
	err = jpeg.Encode(dstFile, img, options)
	if err != nil {
		panic(err)
	}
}
