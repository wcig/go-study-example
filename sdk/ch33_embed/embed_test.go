package ch33_embed

import (
	"embed"
	_ "embed"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// embed
// 嵌入文件到go打包二进制文件中

// 嵌入为[]byte类型 (import _ "embed")
//go:embed hello.txt
var b1 []byte

func TestBytes(t *testing.T) {
	fmt.Println(string(b1)) // hello world!
}

// 嵌入为string类型 (import _ "embed")
//go:embed hello.txt
var s1 string

func TestString(t *testing.T) {
	fmt.Println(s1) // hello world!
}

// 嵌入为文件系统fs (import "embed")
//go:embed hello.txt
var f1 embed.FS

func TestFs(t *testing.T) {
	b, err := f1.ReadFile("hello.txt")
	assert.Nil(t, err)
	fmt.Println(string(b)) // hello world!
}

// 注意: 不支持嵌入为函数内局部变量
func TestLocalVariable(t *testing.T) {
	// //go:embed hello.txt
	// var b []byte
	//
	// //go:embed hello.txt
	// var s string
	//
	// //go:embed hello.txt
	// var f embed.FS

	// output:
	// ./embed_test.go:44:4: go:embed cannot apply to var inside func
	// ./embed_test.go:47:4: go:embed cannot apply to var inside func
	// ./embed_test.go:50:4: go:embed cannot apply to var inside func
}

// 支持嵌入为导出和非导出变量
//go:embed hello.txt
var B2 []byte

//go:embed hello.txt
var S2 string

//go:embed hello.txt
var F2 embed.FS

func TestExportedUnexported(t *testing.T) {
	fmt.Println(string(B2))
	fmt.Println(S2)
	b, _ := F2.ReadFile("hello.txt")
	fmt.Println(string(b))
}

// 支持同一个文件嵌入多个变量
//go:embed hello.txt
var s3 string

//go:embed hello.txt
var s4 string

func TestEmbedOneFileToMultiVariable(t *testing.T) {
	fmt.Println(s3)
	fmt.Println(s4)
}

// 嵌入多个文件到fs方式一
//go:embed hello.txt
//go:embed hello2.txt
var f5 embed.FS

func TestEmbedMultiFileToFs1(t *testing.T) {
	b1, _ := f5.ReadFile("hello.txt")
	fmt.Println(string(b1)) // hello world!

	b2, _ := f5.ReadFile("hello2.txt")
	fmt.Println(string(b2)) // hello world!2
}

// 嵌入多个文件到fs方式二
//go:embed hello.txt hello2.txt
var f6 embed.FS

func TestEmbedMultiFileToFs2(t *testing.T) {
	b1, _ := f6.ReadFile("hello.txt")
	fmt.Println(string(b1)) // hello world!

	b2, _ := f6.ReadFile("hello2.txt")
	fmt.Println(string(b2)) // hello world!2
}

// 支持双引号和反引号
//go:embed "hello.txt" `hello2.txt`
var f7 embed.FS

func TestQuote(t *testing.T) {
	b1, _ := f7.ReadFile("hello.txt")
	fmt.Println(string(b1)) // hello world!

	b2, _ := f7.ReadFile("hello2.txt")
	fmt.Println(string(b2)) // hello world!2
}

// 支持目录
//go:embed tmp
var f8 embed.FS

func TestDir(t *testing.T) {
	b1, _ := f8.ReadFile("tmp/hello.txt")
	fmt.Println(string(b1))

	b2, _ := f8.ReadFile("tmp/hello2.txt")
	fmt.Println(string(b2))
}

// func (f FS) ReadDir(name string) ([]fs.DirEntry, error)
// 读取目录方法
//go:embed tmp
var f9 embed.FS

func TestReadDir(t *testing.T) {
	dirs, err := f9.ReadDir("tmp")
	assert.Nil(t, err)

	for _, dir := range dirs {
		fileInfo, err := dir.Info()
		fmt.Printf("%s, %t, %+v, %+v, %s\n", dir.Name(), dir.IsDir(), dir.Type(), fileInfo, err)
	}
	// output:
	// hello.txt, false, ----------, &{name:tmp/hello.txt data:hello world! hash:[117 9 229 189 160 199 98 210 186 199 249 13 117 139 91 34]}, %!s(<nil>)
	// hello2.txt, false, ----------, &{name:tmp/hello2.txt data:hello world!2 hash:[141 156 121 6 142 89 84 140 113 253 171 22 140 214 122 115]}, %!s(<nil>)
}

// func (f FS) Open(name string) (fs.File, error)
// 打开文件
//go:embed hello.txt
var f10 embed.FS

func TestOpen(t *testing.T) {
	file, err := f10.Open("hello.txt")
	assert.Nil(t, err)

	fileInfo, err := file.Stat()
	assert.Nil(t, err)
	fmt.Println(fileInfo) // &{hello.txt hello world! [117 9 229 189 160 199 98 210 186 199 249 13 117 139 91 34]}

	// file.Read()
	// file.Close()
}

// net/http (http://localhost:28080/)
//go:embed tmp
var f11 embed.FS

func TestHttp(t *testing.T) {
	server := http.Server{
		Addr:    "127.0.0.1:28080",
		Handler: nil,
	}
	http.Handle("/", http.FileServer(http.FS(f11)))
	_ = server.ListenAndServe()
}
