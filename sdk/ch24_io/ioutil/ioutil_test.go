package ioutil

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// io/ioutil：使用的io工具函数（从Go1.16开始相同功迁移至包io和os）

// func ReadAll(r io.Reader) ([]byte, error)：从r读取数据直到出现错误或EOF，返回读取的数据，成功调用返回err==nil而不是err=EOF
func TestReadAll(t *testing.T) {
	str := "hello world."
	buf := bytes.NewBufferString(str)
	b, err := ioutil.ReadAll(buf)
	assert.True(t, err == nil)
	assert.True(t, string(b) == str)
}

// func WriteFile(filename string, data []byte, perm fs.FileMode) error：将数据吸入filename的文件，如果文件不存在则使用perm权限创建，否则将在写入前将其截断(清除并重写)但不修改权限。
func TestWriteFile(t *testing.T) {
	str := "hello world."
	filename := "temp"
	err := ioutil.WriteFile(filename, []byte(str), 0777)
	assert.True(t, err == nil)
}

// func ReadFile(filename string) ([]byte, error)：与ReadAll的不同是从filename的文件读取。
func TestReadFile(t *testing.T) {
	str := "hello world."
	filename := "temp"
	b, err := ioutil.ReadFile(filename)
	assert.True(t, err == nil)
	assert.True(t, string(b) == str)
}

// func ReadDir(dirname string) ([]fs.FileInfo, error)：读取dirname的目录，返回目录下所有内容的fs.FileInfo列表
func TestReadDir(t *testing.T) {
	dirname := "../ioutil"
	rangeDir(dirname)
}

func rangeDir(dirname string, prefixes ...string) {
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		panic(err)
	}

	fmt.Println(strings.Join(prefixes, "") + dirname)
	prefixes = append(prefixes, "\t")
	p := strings.Join(prefixes, "")
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			rangeDir(fileInfo.Name(), p)
		} else {
			fmt.Println(p + fileInfo.Name())
		}
	}
}

// func TempDir(dir, pattern string) (name string, err error)：在目录dir创建一个新的临时目录，目录名称采用模式pattern并在末尾添加随机字符串，返回创建的目录名称
func TestTempDir(t *testing.T) {
	name, err := ioutil.TempDir("dir-temp", "example")
	assert.Nil(t, err)
	fmt.Println(name)

	logsDir, err := ioutil.TempDir("dir-temp", "*-logs")
	assert.Nil(t, err)
	fmt.Println(logsDir)
	// output:
	// dir-temp/example203517291
	// dir-temp/216701902-logs
}

// func TempFile(dir, pattern string) (f *os.File, err error)：在目录dir创建指定模式pattern的临时文件，返回文件的*os.File和错误
func TestTempFile(t *testing.T) {
	tempFile, err := ioutil.TempFile("dir-temp", "example")
	assert.Nil(t, err)
	fmt.Println(tempFile.Name())
	tempFile.Close()

	logFile, err := ioutil.TempFile("dir-temp", "log-*")
	assert.Nil(t, err)
	fmt.Println(logFile.Name())
	logFile.Close()
	// output:
	// dir-temp/example255306848
	// dir-temp/log-774969919
}

// func NopCloser(r io.Reader) io.ReadCloser：io.Reader->io.ReadCloser，其中返回的io.ReadCloser的Close方法为无操作的。
func TestNopCloser(t *testing.T) {
	readCloser := ioutil.NopCloser(os.Stdin)
	err := readCloser.Close()
	assert.Nil(t, err)
}
