package md5

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"testing"
)

func BenchmarkStrMD5Cal11(b *testing.B) {
	data := []byte("hello world.")
	for i := 0; i < b.N; i++ {
		sum := md5.Sum(data)
		_ = fmt.Sprintf("%x", sum)
	}
}

func BenchmarkStrMD5Cal12(b *testing.B) {
	data := "hello world."
	for i := 0; i < b.N; i++ {
		sum := md5.Sum([]byte(data))
		_ = fmt.Sprintf("%x", sum)
	}
}

func BenchmarkStrMD5Cal21(b *testing.B) {
	data := []byte("hello world.")
	for i := 0; i < b.N; i++ {
		sum := md5.Sum(data)
		_ = hex.EncodeToString(sum[:])
	}
}

func BenchmarkStrMD5Cal22(b *testing.B) {
	data := "hello world."
	for i := 0; i < b.N; i++ {
		sum := md5.Sum([]byte(data))
		_ = hex.EncodeToString(sum[:])
	}
}

func BenchmarkStrMD5Cal31(b *testing.B) {
	data := []byte("hello world.")
	for i := 0; i < b.N; i++ {
		h := md5.New()
		h.Write(data)
		_ = hex.EncodeToString(h.Sum(nil))
	}
}

func BenchmarkStrMD5Cal32(b *testing.B) {
	data := "hello world."
	for i := 0; i < b.N; i++ {
		h := md5.New()
		h.Write([]byte(data))
		_ = hex.EncodeToString(h.Sum(nil))
	}
}

func BenchmarkStrMD5Cal41(b *testing.B) {
	data := []byte("hello world.")
	for i := 0; i < b.N; i++ {
		h := md5.New()
		io.WriteString(h, string(data))
		_ = hex.EncodeToString(h.Sum(nil))
	}
}

func BenchmarkStrMD5Cal42(b *testing.B) {
	data := "hello world."
	for i := 0; i < b.N; i++ {
		h := md5.New()
		io.WriteString(h, data)
		_ = hex.EncodeToString(h.Sum(nil))
	}
}

func BenchmarkFileMD5Cal1(b *testing.B) {
	// 大文件处理不推荐
	filename := "src.txt"
	for i := 0; i < b.N; i++ {
		b, _ := os.ReadFile(filename)
		sum := md5.Sum(b)
		_ = hex.EncodeToString(sum[:])
	}
}

func BenchmarkFileMD5Cal2(b *testing.B) {
	filename := "src.txt"
	for i := 0; i < b.N; i++ {
		f, _ := os.Open(filename)
		defer f.Close()

		h := md5.New()
		io.Copy(h, f)
		_ = hex.EncodeToString(h.Sum(nil))
	}
}

func BenchmarkFileMD5Cal3(b *testing.B) {
	filename := "src.txt"
	for i := 0; i < b.N; i++ {
		f, _ := os.Open(filename)
		defer f.Close()

		r := bufio.NewReader(f)
		h := md5.New()
		io.Copy(h, r)
		_ = hex.EncodeToString(h.Sum(nil))
	}
}

// ➜  md5 git:(master) ✗ go test -v -bench=. -run=MD5 -benchmem
// goos: darwin
// goarch: amd64
// pkg: go-app/sdk/ch9_crypto/md5
// cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
// BenchmarkStrMD5Cal11
// BenchmarkStrMD5Cal11-8           3116684               370.7 ns/op            64 B/op          3 allocs/op
// BenchmarkStrMD5Cal12
// BenchmarkStrMD5Cal12-8           3209930               374.5 ns/op            64 B/op          3 allocs/op
// BenchmarkStrMD5Cal21
// BenchmarkStrMD5Cal21-8           7781961               151.2 ns/op            32 B/op          1 allocs/op
// BenchmarkStrMD5Cal22
// BenchmarkStrMD5Cal22-8           7637300               157.1 ns/op            32 B/op          1 allocs/op
// BenchmarkStrMD5Cal31
// BenchmarkStrMD5Cal31-8           6378955               175.5 ns/op            48 B/op          2 allocs/op
// BenchmarkStrMD5Cal32
// BenchmarkStrMD5Cal32-8           6590385               180.8 ns/op            48 B/op          2 allocs/op
// BenchmarkStrMD5Cal41
// BenchmarkStrMD5Cal41-8           4517343               264.9 ns/op           176 B/op          5 allocs/op
// BenchmarkStrMD5Cal42
// BenchmarkStrMD5Cal42-8           4827118               244.4 ns/op           160 B/op          4 allocs/op
// BenchmarkFileMD5Cal1
// BenchmarkFileMD5Cal1-8            115054             10208 ns/op             864 B/op          6 allocs/op
// BenchmarkFileMD5Cal2
// BenchmarkFileMD5Cal2-8             55965             20078 ns/op           33074 B/op          7 allocs/op
// BenchmarkFileMD5Cal3
// BenchmarkFileMD5Cal3-8            221420              5987 ns/op            4490 B/op          8 allocs/op
// PASS
// ok      go-app/sdk/ch9_crypto/md5       16.056s
