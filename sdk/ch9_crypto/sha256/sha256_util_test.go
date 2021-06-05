package sha256

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtilSha256(t *testing.T) {
	v1 := StrSha256("hello world.")
	fmt.Println(v1)

	v2, err := FileSha256("src.txt")
	assert.Nil(t, err)
	fmt.Println(v2)
}

func TestUtilSha224(t *testing.T) {
	v1 := StrSha224("hello world.")
	fmt.Println(v1)

	v2, err := FileSha224("src.txt")
	assert.Nil(t, err)
	fmt.Println(v2)
}

/* ------------------------------------------------------------------- */

func StrSha256(data string) string {
	sum := sha256.Sum256([]byte(data))
	return hex.EncodeToString(sum[:])
}

func FileSha256(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	h := sha256.New()
	if _, err = io.Copy(h, r); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func StrSha224(data string) string {
	sum := sha256.Sum224([]byte(data))
	return hex.EncodeToString(sum[:])
}

func FileSha224(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	h := sha256.New224()
	if _, err = io.Copy(h, r); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
