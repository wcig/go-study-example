package bitmap

import (
	"fmt"
	"strings"
)

// bitmap (参考: https://www.cnblogs.com/aiandbigdata/p/11432310.html)
type Bitmap struct {
	bits []byte
	len  uint
	cap  uint
}

func New(size uint) *Bitmap {
	return &Bitmap{
		bits: make([]byte, (size+7)/8),
		len:  0,
		cap:  size,
	}
}

func (b *Bitmap) Set(num uint) bool {
	if !b.checkRange(num) {
		return false
	}

	byteIndex, bitPos := b.offset(num)
	b.bits[byteIndex] |= 1 << bitPos
	b.len++
	return true
}

func (b *Bitmap) Clear(num uint) bool {
	if !b.checkRange(num) {
		return false
	}

	byteIndex, bitPos := b.offset(num)
	if (b.bits[byteIndex] & (1 << bitPos)) == 0 {
		return true
	}

	b.bits[byteIndex] &= ^(1 << bitPos)
	b.len--
	return true
}

func (b *Bitmap) Contains(num uint) bool {
	if !b.checkRange(num) {
		return false
	}

	byteIndex, bitPos := b.offset(num)
	return !((b.bits[byteIndex] & (1 << bitPos)) == 0)
}

func (b *Bitmap) Size() uint {
	return b.len
}

func (b *Bitmap) Capacity() uint {
	return b.cap
}

func (b *Bitmap) IsEmpty() bool {
	return b.len == 0
}

func (b *Bitmap) String() string {
	var sb strings.Builder
	for i := len(b.bits) - 1; i >= 0; i-- {
		sb.WriteString(fmt.Sprintf("%08b", b.bits[i]))
		if i > 0 {
			sb.WriteString("-")
		}
	}
	return sb.String()
}

func (b *Bitmap) checkRange(num uint) bool {
	return num <= b.cap
}

func (b *Bitmap) offset(num uint) (byteIndex, bitPos uint) {
	byteIndex = num / 8
	bitPos = num % 8
	return byteIndex, bitPos
}
