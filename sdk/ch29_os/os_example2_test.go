package ch29_os

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"
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

	// 方式1: os.Stat()
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

	// 方式2: file.Stat()
	f, err := os.Open(filename)
	if err != nil {
		t.Fatal(err)
	}
	fileInfo2, _ := f.Stat()
	fmt.Println(fileInfo2.Name()) // temp
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

// Chtimes
func TestChtimes(t *testing.T) {
	filename := "temp"
	os.Create(filename)
	// defer os.Remove(filename)

	mtime := time.Date(2006, time.February, 1, 3, 4, 5, 0, time.UTC)
	atime := time.Date(2007, time.March, 2, 4, 5, 6, 0, time.UTC)
	if err := os.Chtimes(filename, atime, mtime); err != nil {
		log.Fatal(err)
	}
}

// Environ/Clearenv
func TestEnv(t *testing.T) {
	vals := os.Environ()
	fmt.Println(len(vals), vals[len(vals)-1]) // 41 CGO_ENABLED=1

	os.Clearenv()

	fmt.Println(len(os.Environ())) // 0
}

// Executable
func TestExecutable(t *testing.T) {
	name, err := os.Executable()
	fmt.Println(name, err) // /private/var/folders/vh/lks7z1qx6x90j10nwtm3njlw0000gn/T/___TestExecutable_in_go_app_sdk_ch29_os.test <nil>
}

// Expand
func TestExpand(t *testing.T) {
	mapper := func(placeholderName string) string {
		switch placeholderName {
		case "DAY_PART":
			return "morning"
		case "NAME":
			return "Gopher"
		}

		return ""
	}

	fmt.Println(os.Expand("Good ${DAY_PART}, $NAME!", mapper)) // Good morning, Gopher!
}

// ExpandEnv
func TestExpandEnv(t *testing.T) {
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")

	fmt.Println(os.ExpandEnv("$NAME lives in ${BURROW}."))
}

// LookupEnv
func TestLookupEnv(t *testing.T) {
	val, has := os.LookupEnv("GOROOT")
	fmt.Println(val, has)

	val, has = os.LookupEnv("GOROOT2")
	fmt.Println(val, has)
}

// Unsetenv
func TestUnsetenv(t *testing.T) {
	os.Setenv("TMPDIR", "/my/tmp")
	fmt.Println(os.Getenv("TMPDIR")) // /my/tmp

	os.Unsetenv("TMPDIR")

	fmt.Println(os.Getenv("TMPDIR")) //
}

// SameFile
func TestSameFile(t *testing.T) {
	f1, _ := os.Stat("/bin/bzcmp")
	f2, _ := os.Stat("/bin/bzdiff")
	fmt.Println(os.SameFile(f1, f2)) // true
}

// link
func TestLinkAndSymlink(t *testing.T) {
	f1, _ := os.Create("temp1")
	os.Link("temp1", "temp2")
	f1.Close()

	f3, _ := os.Create("temp3")
	os.Symlink("temp3", "temp4")
	f3.Close()
}

// TempDir
func TestTempDir(t *testing.T) {
	fmt.Println(os.TempDir()) // tmp
}

// UserDir
func TestUserDir(t *testing.T) {
	fmt.Println(os.UserCacheDir())  // /home/yangbo/.cache <nil>
	fmt.Println(os.UserConfigDir()) // /home/yangbo/.config <nil>
	fmt.Println(os.UserHomeDir())   // /home/yangbo <nil>
}

// Get
func TestGet(t *testing.T) {
	fmt.Println(os.Geteuid(), os.Getuid()) // 501 501
	fmt.Println(os.Getegid(), os.Getgid()) // 20 20
	fmt.Println(os.Getgroups())            // [20 12 61 79 80 81 98 33 100 204 250 395 398 399 400 701] <nil>
	fmt.Println(os.Getpagesize())          // 4096
	fmt.Println(os.Getpid())               // 8785
	fmt.Println(os.Getppid())              // 8784
	fmt.Println(os.Getwd())                // /Users/yangbo/Documents/MyGithub/go-study-example/ch12_io/os <nil>
	fmt.Println(os.Hostname())             // ubuntu <nil>
}

// File
func TestFile(t *testing.T) {
	f1 := os.Stdin
	fmt.Println(f1.Fd())
	fmt.Println(f1.Name())

	f2, _ := os.Create("temp")
	fmt.Println(f2.Fd())
	fmt.Println(f2.Name())

	f2.Close()
	os.Remove("temp")

	f3, _ := os.Open("../../ch12_io")
	f3Infos, _ := f3.Readdir(5)
	for _, fileInfo := range f3Infos {
		fmt.Println(fileInfo.Name())
	}

	f4, _ := os.Open("../../ch12_io")
	f4Names, _ := f4.Readdirnames(-1)
	for _, name := range f4Names {
		fmt.Println(name)
	}
}
