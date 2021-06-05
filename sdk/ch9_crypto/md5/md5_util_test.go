package md5

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtil(t *testing.T) {
	v1 := StrMD5("hello world.")
	fmt.Println(v1)

	v2, err := FileMD5("src.txt")
	assert.Nil(t, err)
	fmt.Println(v2)
}

/* ------------------------------------------------------------------- */

// 计算字符串MD5
func StrMD5(data string) string {
	sum := md5.Sum([]byte(data))
	return hex.EncodeToString(sum[:])
}

// 计算文件MD5
func FileMD5(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	h := md5.New()
	if _, err = io.Copy(h, r); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
