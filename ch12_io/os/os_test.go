package os

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// FileMode
func TestGetFileMode(t *testing.T) {
	filename := "temp"
	os.Create(filename)
	defer os.Remove(filename)

	fileInfo, err := os.Stat(filename)
	if err != nil {
		t.Fatal(err)
	}
	fileMode := fileInfo.Mode()
	fmt.Println(fileMode.IsDir(), fileMode.IsRegular(), fileMode.Perm(), fileMode.String()) // false true -rw-rw-r-- -rw-rw-r--
}

// FileInfo
func TestGetFileInfo(t *testing.T) {
	filename := "temp"
	os.Create(filename)
	defer os.Remove(filename)

	fileInfo, err := os.Stat(filename)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fileInfo.IsDir())   // false
	fmt.Println(fileInfo.Mode())    // -rw-rw-r--
	fmt.Println(fileInfo.Name())    // temp
	fmt.Println(fileInfo.Size())    // 0
	fmt.Println(fileInfo.ModTime()) // 2020-09-07 19:46:26.432748332 +0800 CST
	fmt.Println(fileInfo.Sys())     // &{2050 5113192 1 33204 1000 1000 0 0 0 4096 0 {1599479186 432748332} {1599479186 432748332} {1599479186 432748332} [0 0 0]}

	// 相对路径 -> 绝对路径
	path, err := filepath.Abs(fileInfo.Name())
	fmt.Println(path) // /home/yangbo/Documents/workspace/go-study-example/ch12_io/os/temp
}

// Chmod
func TestChmod(t *testing.T) {
	filename := "temp"
	os.Create(filename)
	defer os.Remove(filename)

	fileInfo, _ := os.Stat(filename)
	fmt.Println(fileInfo.Mode()) // -rw-rw-r--

	newMode := os.FileMode(0777) // 权限格式: 0xxx
	os.Chmod(filename, newMode)

	fileInfo, _ = os.Stat(filename)
	fmt.Println(fileInfo.Mode()) // -rwxrwxrwx
}
