package graphicsmagick

import (
	"fmt"
	"testing"

	"github.com/gographics/gmagick"
)

// graphicsmagick
// version: v1.3.28, module: github.com/gographics/gmagick v1.0.0

// 获取视频宽高
func TestGetImageSize(t *testing.T) {
	gmagick.Initialize()
	defer gmagick.Terminate()

	mw := gmagick.NewMagickWand()
	defer mw.Destroy()

	var err error
	if err = mw.ReadImage("src.jpeg"); err != nil {
		panic(err)
	}

	w := mw.GetImageWidth()
	h := mw.GetImageHeight()
	fmt.Println("image weight and height:", w, h) // image weight and height: 360 640
}

// 获取视频宽高
func TestGetImageSize2(t *testing.T) {
	gmagick.Initialize()
	defer gmagick.Terminate()

	mw := gmagick.NewMagickWand()
	defer mw.Destroy()

	var err error
	if err = mw.ReadImage("src.jpeg"); err != nil {
		panic(err)
	}

	w, h, err := mw.GetSize()
	if err != nil {
		panic(err)
	}
	fmt.Println("image weight and height:", w, h) // image weight and height: 0 0
}

// 图普调整大小
func TestResize(t *testing.T) {
	gmagick.Initialize()
	defer gmagick.Terminate()

	mw := gmagick.NewMagickWand()
	defer mw.Destroy()

	var err error
	if err = mw.ReadImage("src.jpeg"); err != nil {
		panic(err)
	}

	filter := gmagick.FILTER_LANCZOS
	w := mw.GetImageWidth()
	h := mw.GetImageHeight()
	err = mw.ResizeImage(w/2, h/2, filter, 1)
	if err != nil {
		panic(err)
	}

	err = mw.WriteImage("tmp.jpeg")
	if err != nil {
		panic(err)
	}
}

// func resize(orig string, dest string) {
//    mw := gmagick.NewMagickWand()
//    defer mw.Destroy()
//    mw.ReadImage(orig)
//    filter := gmagick.FILTER_LANCZOS
//    w := mw.GetImageWidth()
//    h := mw.GetImageHeight()
//    mw.ResizeImage(w/2, h/2, filter, 1)
//    mw.WriteImage(dest)
// }
//
// func main() {
//    f := flag.String("from", "", "original image file ...")
//    t := flag.String("to", "", "target file ...")
//    flag.Parse()
//
//    gmagick.Initialize()
//    defer gmagick.Terminate()
//
//    resize(*f, *t)
// }
