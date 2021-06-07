package base32

import (
	"encoding/base32"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtil(t *testing.T) {
	src := "hello world."
	ecStr := Base32EncodeToString([]byte(src))
	fmt.Println("base32 encode result:", ecStr)

	dcBytes, err := Base32DecodeString(ecStr)
	assert.Nil(t, err)
	fmt.Println("base32 decode result:", string(dcBytes))
	// output:
	// base32 encode result: NBSWY3DPEB3W64TMMQXA====
	// base32 decode result: hello world.
}

/* ------------------------------------------------------------------- */

func Base32EncodeToString(src []byte) string {
	return base32.StdEncoding.EncodeToString(src)
}

func Base32DecodeString(s string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(s)
}
