package main

import (
	"image/color"
	"os"
	"testing"

	"github.com/skip2/go-qrcode"
)

func TestFirst(t *testing.T) {
	var png []byte
	png, err := qrcode.Encode("https://example.org", qrcode.Medium, 256)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("qr.png", png, os.FileMode(0644))
	if err != nil {
		panic(err)
	}
}

func TestWriteFile(t *testing.T) {
	err := qrcode.WriteFile("https://example.org", qrcode.Medium, 256, "qr.png")
	if err != nil {
		panic(err)
	}
}

func TestCustomColors(t *testing.T) {
	err := qrcode.WriteColorFile("https://example.org", qrcode.Medium, 256,
		color.Black, color.White, "qr.png")
	if err != nil {
		panic(err)
	}
}
