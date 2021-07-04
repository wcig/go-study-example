package png

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	srcFile, err := os.Open("src.png")
	assert.Nil(t, err)

	imgCfg, err := png.DecodeConfig(srcFile)
	assert.Nil(t, err)
	fmt.Printf("image width and height: %dx%d\n", imgCfg.Width, imgCfg.Height)
	srcFile.Close()

	srcFile, err = os.Open("src.png")
	assert.Nil(t, err)
	defer srcFile.Close()
	img, err := png.Decode(srcFile)
	assert.Nil(t, err)

	dstFile, err := os.Create("tmp.png")
	assert.Nil(t, err)
	defer dstFile.Close()

	var ec png.Encoder
	ec.CompressionLevel = png.BestCompression
	err = ec.Encode(dstFile, img)
	assert.Nil(t, err)
}

func TestEncode(t *testing.T) {
	const width, height = 256, 256

	// Create a colored image of the given width and height.
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.NRGBA{
				R: uint8((x + y) & 255),
				G: uint8((x + y) << 1 & 255),
				B: uint8((x + y) << 2 & 255),
				A: 255,
			})
		}
	}

	f, err := os.Create("tmp.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
