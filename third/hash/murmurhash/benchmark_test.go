package murmurhash

import (
	"crypto/md5"
	"crypto/sha1"
	"testing"

	"github.com/spaolacci/murmur3"
)

var data = []byte("hello world")

func md5Hash() [16]byte {
	return md5.Sum(data)
}
func sha1Hash() [20]byte {
	return sha1.Sum(data)
}
func murmur32() uint32 {
	return murmur3.Sum32(data)
}
func murmur64() uint64 {
	return murmur3.Sum64(data)
}

func BenchmarkMD5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		md5Hash()
	}
}
func BenchmarkSHA1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sha1Hash()
	}
}
func BenchmarkMurmurHash32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		murmur32()
	}
}
func BenchmarkMurmurHash64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		murmur64()
	}
}

// Output:
// goos: darwin
// goarch: arm64
// pkg: go-app/third/hash/murmurhash
// BenchmarkMD5
// BenchmarkMD5-8            	10913847	       108.3 ns/op
// BenchmarkSHA1
// BenchmarkSHA1-8           	21559855	        55.32 ns/op
// BenchmarkMurmurHash32
// BenchmarkMurmurHash32-8   	223802968	         5.320 ns/op
// BenchmarkMurmurHash64
// BenchmarkMurmurHash64-8   	131596021	         9.068 ns/op
