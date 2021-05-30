package imagemagick

import (
	"fmt"
	"testing"

	"gopkg.in/gographics/imagick.v2/imagick"
)

func Test(t *testing.T) {
	// imagick.Initialize()
	// defer imagick.Terminate()

	// mw := imagick.NewMagickWand()
	// mw.Destroy()

	imageCommandResult, err := imagick.ConvertImageCommand([]string{
		"convert", "/Users/yangbo/imagemagick/test.png", "-resize", "200x200", "/Users/yangbo/imagemagick/test-out.png",
	})
	fmt.Println(err)                // <nil>
	fmt.Println(imageCommandResult) // &{0xc000010040 200,200,PNG}
}
