package textproto

import (
	"fmt"
	"net/textproto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrimBytes(t *testing.T) {
	b := textproto.TrimBytes([]byte(" abc  "))
	assert.Equal(t, "abc", string(b))
}

func TestTrimString(t *testing.T) {
	b := textproto.TrimString(" abc  ")
	assert.Equal(t, "abc", b)
}

func TestMIMEHeader(t *testing.T) {
	h := textproto.MIMEHeader{}
	h.Add("a", "a1")
	h.Add("a", "a2")
	fmt.Println(h)

	v1 := h.Get("a")
	fmt.Println(v1)

	v2 := h.Values("a")
	fmt.Println(v2)

	h.Set("a", "1")
	v3 := h.Values("a")
	fmt.Println(v3)

	h.Del("a")
	fmt.Println(h)
	// output:
	// map[A:[a1 a2]]
	// a1
	// [a1 a2]
	// [1]
	// map[]
}
