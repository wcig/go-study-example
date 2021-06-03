package imagemagick

import (
	"fmt"
	"testing"

	"gopkg.in/gographics/imagick.v2/imagick"
)

// imagemagick
// version: 6.9.7, module: gopkg.in/gographics/imagick.v2 v2.6.0

func TestAdaptiveResizeImage(t *testing.T) {
	var err error

	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

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
