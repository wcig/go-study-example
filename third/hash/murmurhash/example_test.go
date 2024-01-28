package murmurhash

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/spaolacci/murmur3"
)

func TestExample(t *testing.T) {
	in := []byte("hello world")

	sum32 := murmur3.Sum32(in)
	fmt.Println("sum32:", sum32)

	sum64 := murmur3.Sum64(in)
	fmt.Println("sum64:", sum64)

	h1, h2 := murmur3.Sum128(in)
	fmt.Println("sum128:", h1, h2)
	// Output:
	// sum32: 1586663183
	// sum64: 5998619086395760910
	// sum128: 5998619086395760910 12364428806279881649
}

func TestDistribution(t *testing.T) {
	const bucketSize = 10
	var bucketMap = map[uint64]int{}
	for i := 15000000000; i < 15000000000+10000000; i++ {
		val := []byte(strconv.Itoa(i))
		hashInt := murmur3.Sum64(val) % uint64(bucketSize)
		bucketMap[hashInt]++
	}
	fmt.Println(bucketMap)
	// Output:
	// map[0:998000 1:999862 2:1000388 3:999393 4:1000657 5:1000487 6:999575 7:999594 8:1000520 9:1001524]
}
