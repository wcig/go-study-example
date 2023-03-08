package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/tuotoo/qrcode"
)

// "github.com/tuotoo/qrcode"库只能解析特定二维码图片
func TestFirst(t *testing.T) {
	// 打开二维码图片文件
	file, err := os.Open("qr.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 解析二维码图片
	qrmatrix, err := qrcode.Decode(file)
	if err != nil {
		panic(err)
	}

	// 输出二维码中包含的数据
	fmt.Println(qrmatrix.Content)
}

// zbar cli (zbarimg qr.png)
func TestZbraimg(t *testing.T) {
	cmd := exec.Command("zbarimg", "-q", "qr.png")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	result := strings.TrimPrefix(out.String(), "QR-Code:")
	fmt.Printf(result)
}
