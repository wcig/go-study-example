package bitmap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitmap(t *testing.T) {
	const size = 16
	bm := New(size)
	fmt.Println(">> after init")
	printBitmap(bm)

	// set
	var expectArr []int
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			bm.Set(uint(i))
			expectArr = append(expectArr, i)
		}
	}
	assert.Equal(t, uint(size/2), bm.Size())
	assert.Equal(t, false, bm.IsEmpty())
	fmt.Println(">> after set")
	printBitmap(bm)

	// contain
	var testArr []int
	for i := 0; i < size; i++ {
		if bm.Contains(uint(i)) {
			testArr = append(testArr, i)
		}
	}
	assert.Equal(t, expectArr, testArr)

	// clear
	for i := 0; i < size; i++ {
		bm.Clear(uint(i))
	}
	assert.Equal(t, uint(0), bm.Size())
	assert.Equal(t, true, bm.IsEmpty())
	fmt.Println(">> after clear")
	printBitmap(bm)

	testArr = []int{}
	for i := 0; i < size; i++ {
		if bm.Contains(uint(i)) {
			testArr = append(testArr, i)
		}
	}
	assert.Equal(t, 0, len(testArr))

	// Output:
	// >> after init
	// bitmap cap: 16, size: 0, empty: true, value: 00000000-00000000
	// >> after set
	// bitmap cap: 16, size: 8, empty: false, value: 01010101-01010101
	// >> after clear
	// bitmap cap: 16, size: 0, empty: true, value: 00000000-00000000
}

func printBitmap(bm *Bitmap) {
	fmt.Printf("bitmap cap: %d, size: %d, empty: %t, value: %s\n", bm.Capacity(), bm.Size(), bm.IsEmpty(), bm)
}
