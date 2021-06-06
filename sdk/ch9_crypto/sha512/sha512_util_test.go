package sha512

import (
	"bufio"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtilSha512(t *testing.T) {
	v1 := StrSha512("hello world.")
	fmt.Println(v1)

	v2, err := FileSha512("src.txt")
	assert.Nil(t, err)
	fmt.Println(v2)
}

/* ------------------------------------------------------------------- */

func StrSha512(data string) string {
	sum := sha512.Sum512([]byte(data))
	return hex.EncodeToString(sum[:])
}

func FileSha512(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	h := sha512.New()
	if _, err = io.Copy(h, r); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
