package base64

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtil(t *testing.T) {
	src := "hello world."
	ecStr := Base64EncodeToString([]byte(src))
	fmt.Println("base64 encode result:", ecStr)

	dcBytes, err := Base64DecodeString(ecStr)
	assert.Nil(t, err)
	fmt.Println("base64 decode result:", string(dcBytes))
	// output:
	// base64 encode result: aGVsbG8gd29ybGQu
	// base64 decode result: hello world.
}

/* ------------------------------------------------------------------- */

func Base64EncodeToString(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func Base64DecodeString(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}
