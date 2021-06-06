package sha1

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtilSha1(t *testing.T) {
	v1 := StrSha1("hello world.")
	fmt.Println(v1)

	v2, err := FileSha1("src.txt")
	assert.Nil(t, err)
	fmt.Println(v2)
}

/* ------------------------------------------------------------------- */

func StrSha1(data string) string {
	sum := sha1.Sum([]byte(data))
	return hex.EncodeToString(sum[:])
}

func FileSha1(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	h := sha1.New()
	if _, err = io.Copy(h, r); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
