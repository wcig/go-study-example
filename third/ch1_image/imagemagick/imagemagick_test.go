package imagemagick

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/gographics/imagick.v2/imagick"
)

// imagemagick
// version: 6.9.7, module: gopkg.in/gographics/imagick.v2 v2.6.0

func TestGetImageSize(t *testing.T) {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	var err error
	if err = mw.ReadImage("src.jpeg"); err != nil {
		panic(err)
	}

	fmt.Printf("image w:%d, h:%d\n", mw.GetImageWidth(), mw.GetImageHeight()) // image w:360, h:640
}

func TestAdaptiveResizeImage(t *testing.T) {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	var err error
	if err = mw.ReadImage("src.jpeg"); err != nil {
		panic(err)
	}

	if err = mw.AdaptiveResizeImage(200, 200); err != nil {
		panic(err)
	}

	if err = mw.WriteImage("tmp.200x200.jpeg"); err != nil {
		panic(err)
	}
}

func TestConvertCmd(t *testing.T) {
	imageCommandResult, err := imagick.ConvertImageCommand([]string{
		"convert", "src.jpeg", "-resize", "200x200", "tmp.200x200.jpeg",
	})
	fmt.Println(err)                // <nil>
	fmt.Println(imageCommandResult) // &{0xc000010040 113,200,JPEG}
}

// 调整尺寸并居中裁剪图片
// convert input.png -resize "1080x1080>" -gravity center -extent "1080x1080>" output_1080x1080.png
func TestResizeAndGravityCenter(t *testing.T) {
	imagick.Initialize()
	defer imagick.Terminate()

	mw1 := imagick.NewMagickWand()
	defer mw1.Destroy()

	var err error
	err = mw1.ReadImage("src.jpg")
	assert.Nil(t, err)

	var w, h uint = 1080, 1080
	size := fmt.Sprintf("%dx%d^+0+0", w, h)
	mw2 := mw1.TransformImage("", size)
	defer mw2.Destroy()

	err = mw2.SetImageGravity(imagick.GRAVITY_CENTER)
	assert.Nil(t, err)

	offsetX := -(int(w) - int(mw2.GetImageWidth())) / 2
	offsetY := -(int(h) - int(mw2.GetImageHeight())) / 2
	err = mw2.ExtentImage(w, h, offsetX, offsetY)
	assert.Nil(t, err)

	err = mw2.WriteImage("tmp.1080x1080.jpg")
	assert.Nil(t, err)
}
