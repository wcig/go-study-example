package ch30_path

import (
	"fmt"
	"path"
	"testing"
)

// path: 实现了用于操作斜杠分隔路径的实用程序。
// path 包应该只用于由正斜杠分隔的路径，例如 URL 中的路径。这个包不处理带有驱动器号或反斜杠的 Windows 路径；要操作操作系统路径，请使用 path/filepath 包。

// 错误
func TestErr(t *testing.T) {
	_ = path.ErrBadPattern // syntax error in pattern
}

// func Base(path string) string: 返回路径的最后一个元素，去除尾部斜杠（空路径则返回'.'）
func TestBase(t *testing.T) {
	fmt.Println(path.Base("/a/b/c.txt"))
	fmt.Println(path.Base("a/b/c.txt"))
	fmt.Println(path.Base("/"))
	fmt.Println(path.Base(""))
	// output:
	// c.txt
	// c.txt
	// /
	// .
}

// func Clean(path string) string: 返回短路径
func TestClean(t *testing.T) {
	paths := []string{
		"a/c",
		"a//c",
		"a/c/.",
		"a/c/b/..",
		"/../a/c",
		"/../a/b/../././/c",
		"",
	}

	for _, p := range paths {
		fmt.Printf("Clean(%q) = %q\n", p, path.Clean(p))
	}
	// output:
	// Clean("a/c") = "a/c"
	// Clean("a//c") = "a/c"
	// Clean("a/c/.") = "a/c"
	// Clean("a/c/b/..") = "a/c"
	// Clean("/../a/c") = "/a/c"
	// Clean("/../a/b/../././/c") = "/a/c"
	// Clean("") = "."
}

// func Dir(path string) string: 返回除最后一个元素之外的路径（空路径则返回'.'）
func TestDir(t *testing.T) {
	fmt.Println(path.Dir("/a/b/c"))
	fmt.Println(path.Dir("a/b/c"))
	fmt.Println(path.Dir("/a/"))
	fmt.Println(path.Dir("a/"))
	fmt.Println(path.Dir("/"))
	fmt.Println(path.Dir(""))
	// output:
	// /a/b
	// a/b
	// /a
	// a
	// /
	// .
}

// func Ext(path string) string: 返回文件扩展名，包括'.'符号（没有则返回空）
func TestExt(t *testing.T) {
	fmt.Println(path.Ext("/a/b/c/bar.css"))
	fmt.Println(path.Ext("/"))
	fmt.Println(path.Ext(""))
	// output:
	// .css
	//
	//
}

// func IsAbs(path string) bool: 报告path是否为绝对路径
func TestIsAbs(t *testing.T) {
	fmt.Println(path.IsAbs("/dev/null")) // true
}

// func Join(elem ...string) string: 连接多个元素成路径，用斜线分割（空元素会被忽略，列表为空或所有元素为空字符串则返回空）
func TestJoin(t *testing.T) {
	fmt.Println(path.Join("a", "b", "c"))
	fmt.Println(path.Join("a", "b/c"))
	fmt.Println(path.Join("a/b", "c"))

	fmt.Println(path.Join("a/b", "../../../xyz"))

	fmt.Println(path.Join("", ""))
	fmt.Println(path.Join("a", ""))
	fmt.Println(path.Join("", "a"))
	// output:
	// a/b/c
	// a/b/c
	// a/b/c
	// ../xyz
	//
	// a
	// a
}

// func Match(pattern, name string) (matched bool, err error): 报告name是否与模式patter匹配
func TestMatch(t *testing.T) {
	fmt.Println(path.Match("abc", "abc"))
	fmt.Println(path.Match("a*", "abc"))
	fmt.Println(path.Match("a*/b", "a/c/b"))
	// output:
	// true <nil>
	// true <nil>
	// false <nil>
}

// func Split(path string) (dir, file string): 拆分path为目录和文件名（返回的值具有 path = dir+file 的属性。）
func TestSplit(t *testing.T) {
	split := func(s string) {
		dir, file := path.Split(s)
		fmt.Printf("path.Split(%q) = dir: %q, file: %q\n", s, dir, file)
	}
	split("static/myfile.css")
	split("myfile.css")
	split("")
	// output:
	// path.Split("static/myfile.css") = dir: "static/", file: "myfile.css"
	// path.Split("myfile.css") = dir: "", file: "myfile.css"
	// path.Split("") = dir: "", file: ""
}
